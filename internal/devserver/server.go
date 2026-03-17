package devserver

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//go:embed dashboard.html
var dashboardHTML []byte

// Event is a wrapper around an update with metadata for the dashboard.
type Event struct {
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

// Server is a dev dashboard that broadcasts updates via SSE.
type Server struct {
	clients map[chan []byte]struct{}
	mu      sync.RWMutex
}

// New creates a new dev server.
func New() *Server {
	return &Server{
		clients: make(map[chan []byte]struct{}),
	}
}

// Start starts the dashboard HTTP server on the given address.
func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleDashboard)
	mux.HandleFunc("/events", s.handleSSE)

	fmt.Printf("\033[35m[gramkit:watch]\033[0m Dashboard: http://localhost%s\n", addr)
	return http.ListenAndServe(addr, mux)
}

// Broadcast sends an update to all connected SSE clients.
func (s *Server) Broadcast(data []byte) {
	event, _ := json.Marshal(Event{
		Timestamp: time.Now(),
		Data:      data,
	})

	s.mu.RLock()
	defer s.mu.RUnlock()
	for ch := range s.clients {
		select {
		case ch <- event:
		default:
			// client too slow, drop event
		}
	}
}

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(dashboardHTML)
}

func (s *Server) handleSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ch := make(chan []byte, 64)
	s.mu.Lock()
	s.clients[ch] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, ch)
		s.mu.Unlock()
	}()

	// Send initial connected event.
	fmt.Fprintf(w, "data: {\"type\":\"connected\"}\n\n")
	flusher.Flush()

	for {
		select {
		case data := <-ch:
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

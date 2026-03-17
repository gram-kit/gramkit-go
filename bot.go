package gramkit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gram-kit/gramkit-go/internal/devserver"
	"github.com/gram-kit/gramkit-go/internal/network"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

// HandlerFunc is a function that handles an incoming update.
type HandlerFunc func(ctx context.Context, b *Bot, update *models.Update)

// Bot represents a Telegram bot instance.
type Bot struct {
	client         *network.Client
	defaultHandler HandlerFunc
	handlers       []handler
	middlewares    []MiddlewareFunc
	pollTimeout    int
	devServer      *devserver.Server
	mu             sync.RWMutex
}

// Option is a functional option for configuring the Bot.
type Option func(*Bot)

// WithDefaultHandler sets the default handler for all updates.
func WithDefaultHandler(handler HandlerFunc) Option {
	return func(b *Bot) {
		b.defaultHandler = handler
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(b *Bot) {
		b.client.HTTPClient = client
	}
}

// WithServerURL sets a custom Bot API server URL (e.g. for local Bot API server).
func WithServerURL(url string) Option {
	return func(b *Bot) {
		b.client.BaseURL = fmt.Sprintf("%s/bot%s", url, b.client.Token)
	}
}

// WithPollTimeout sets the long polling timeout in seconds (default: 30).
func WithPollTimeout(seconds int) Option {
	return func(b *Bot) {
		b.pollTimeout = seconds
	}
}

// New creates a new Bot instance with the given token and options.
func New(token string, opts ...Option) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("gramkit: token is empty")
	}

	b := &Bot{
		client:      network.NewClient(token, nil),
		pollTimeout: 30,
	}

	for _, opt := range opts {
		opt(b)
	}

	// Start dev watch dashboard if GRAMKIT_WATCH_ADDR is set.
	if addr := os.Getenv("GRAMKIT_WATCH_ADDR"); addr != "" {
		b.devServer = devserver.New()
		go func() {
			if err := b.devServer.Start(addr); err != nil {
				log.Printf("gramkit: dev server error: %v", err)
			}
		}()
	}

	return b, nil
}

// StartPolling starts long polling. It blocks until the context is cancelled.
func (b *Bot) StartPolling(ctx context.Context, opts ...params.GetUpdates) error {
	var p params.GetUpdates
	if len(opts) > 0 {
		p = opts[0]
	}
	if p.Timeout == 0 {
		p.Timeout = b.pollTimeout
	}

	offset := p.Offset

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		p.Offset = offset
		upds, err := b.GetUpdates(ctx, p)
		if err != nil {
			log.Printf("gramkit: polling error: %v", err)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(3 * time.Second):
				continue
			}
		}

		for _, upd := range upds {
			b.processUpdate(ctx, &upd)
			offset = upd.UpdateID + 1
		}
	}
}

// StartWebhook starts an HTTP server to receive webhook updates. It blocks until the context is cancelled.
func (b *Bot) StartWebhook(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", b.webhookHandler())

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// WebhookHandler returns an http.HandlerFunc for use with your own HTTP server.
func (b *Bot) WebhookHandler() http.HandlerFunc {
	return b.webhookHandler()
}

func (b *Bot) webhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var upd models.Update
		if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b.processUpdate(r.Context(), &upd)
		w.WriteHeader(http.StatusOK)
	}
}

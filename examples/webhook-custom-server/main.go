package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := gramkit.New(os.Getenv("BOT_TOKEN"),
		gramkit.WithDefaultHandler(echoHandler),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Build a custom HTTP mux with the bot webhook alongside other routes.
	mux := http.NewServeMux()

	// Mount the bot webhook handler on a specific path.
	mux.HandleFunc("/webhook", bot.WebhookHandler())

	// Other application routes can live alongside the bot.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Graceful shutdown on context cancellation.
	go func() {
		<-ctx.Done()
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		srv.Shutdown(shutdownCtx)
	}()

	fmt.Println("Custom server listening on :8080")
	fmt.Println("  POST /webhook  - Telegram webhook")
	fmt.Println("  GET  /health   - Health check")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func echoHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	b.SendMessage(ctx, params.SendMessage{
		ChatID: update.Message.Chat.ID,
		Text:   "Echo: " + update.Message.Text,
	})
}

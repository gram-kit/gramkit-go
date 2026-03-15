package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

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

	// Set the webhook URL via the Telegram Bot API before running this bot,
	// for example:
	//   curl -F "url=https://example.com/" https://api.telegram.org/bot<TOKEN>/setWebhook
	//
	// Then start this server so it receives the incoming updates.

	// Listen for incoming webhook updates on port 8080.
	addr := ":8080"
	fmt.Printf("Listening for webhooks on %s\n", addr)
	if err := bot.StartWebhook(ctx, addr); err != nil {
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

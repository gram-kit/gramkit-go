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
		gramkit.WithDefaultHandler(defaultHandler),
		gramkit.WithMiddleware(loggerMiddleware),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register command handlers (typed — receives *models.Message directly)
	bot.HandleCommand("start", startHandler)
	bot.HandleCommand("help", helpHandler)

	// Register callback query handler with prefix match (typed — receives *models.CallbackQuery directly)
	bot.HandleCallbackQueryMatch(gramkit.MatchPrefix, "action_", callbackHandler)

	fmt.Println("Bot is running...")
	bot.StartPolling(ctx)
}

// loggerMiddleware logs every incoming update.
func loggerMiddleware(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)
	}
	next(ctx, b, update)
}

func startHandler(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID:    msg.Chat.ID,
		Text:      "Hello! Powered by -gramkit-",
		ParseMode: "Markdown",
	})
}

func helpHandler(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "Commands:\n/start - start\n/help - help\nAny text - echo",
	})
}

func callbackHandler(ctx context.Context, b *gramkit.Bot, cq *models.CallbackQuery) {
	log.Printf("Callback: %s", cq.Data)
}

// defaultHandler handles all unmatched messages (still uses *models.Update for WithDefaultHandler).
func defaultHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, params.SendMessage{
			ChatID: update.Message.Chat.ID,
			Text:   "Echo: " + update.Message.Text,
		})
	}
}

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

	// Register specific command handlers
	bot.RegisterHandler(gramkit.OnMessage, "/start", startHandler)
	bot.RegisterHandler(gramkit.OnMessage, "/help", helpHandler)

	// Register callback query handler with prefix match
	bot.RegisterHandlerMatch(gramkit.OnCallbackQuery, gramkit.MatchPrefix, "action_", callbackHandler)

	fmt.Println("Бот запущен...")
	bot.StartPolling(ctx)
}

// loggerMiddleware logs every incoming update.
func loggerMiddleware(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)
	}
	next(ctx, b, update)
}

func startHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет! Я echo-бот на gramkit 🤖",
	})
}

func helpHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: update.Message.Chat.ID,
		Text:   "Команды:\n/start - начало\n/help - помощь\nЛюбой текст - эхо",
	})
}

func callbackHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	log.Printf("Callback: %s", update.CallbackQuery.Data)
}

// defaultHandler handles all unmatched messages.
func defaultHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, params.SendMessage{
			ChatID: update.Message.Chat.ID,
			Text:   "Эхо: " + update.Message.Text,
		})
	}
}

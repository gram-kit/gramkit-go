package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Parse allowed user IDs from a comma-separated env var, e.g. "123456,789012".
	allowedIDs := parseAllowedUsers(os.Getenv("ALLOWED_USER_IDS"))

	bot, err := gramkit.New(os.Getenv("BOT_TOKEN"),
		gramkit.WithDefaultHandler(defaultHandler),
		// Middleware is executed in order: logger first, then auth check.
		gramkit.WithMiddleware(
			loggerMiddleware,
			authMiddleware(allowedIDs),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	bot.RegisterHandler(gramkit.OnMessage, "/start", startHandler)

	fmt.Println("Middleware bot is running...")
	bot.StartPolling(ctx)
}

// loggerMiddleware logs every incoming update and passes it along.
func loggerMiddleware(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	if update.Message != nil {
		log.Printf("[LOG] from=%s text=%q", update.Message.From.FirstName, update.Message.Text)
	}
	next(ctx, b, update)
}

// authMiddleware returns a middleware that only allows updates from the given
// user IDs. If the set is empty, all users are allowed.
func authMiddleware(allowed map[int64]struct{}) gramkit.MiddlewareFunc {
	return func(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
		// If no whitelist is configured, allow everyone.
		if len(allowed) == 0 {
			next(ctx, b, update)
			return
		}

		var userID int64
		switch {
		case update.Message != nil && update.Message.From != nil:
			userID = update.Message.From.ID
		case update.CallbackQuery != nil:
			userID = update.CallbackQuery.From.ID
		default:
			// Unknown update type -- let it through.
			next(ctx, b, update)
			return
		}

		if _, ok := allowed[userID]; !ok {
			log.Printf("[AUTH] blocked user %d", userID)
			// Do NOT call next() -- this effectively blocks the update.
			return
		}

		next(ctx, b, update)
	}
}

func startHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome! You are authorized.",
	})
}

func defaultHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, params.SendMessage{
			ChatID: update.Message.Chat.ID,
			Text:   "Echo: " + update.Message.Text,
		})
	}
}

func parseAllowedUsers(raw string) map[int64]struct{} {
	result := make(map[int64]struct{})
	if raw == "" {
		return result
	}
	for _, s := range strings.Split(raw, ",") {
		s = strings.TrimSpace(s)
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		result[id] = struct{}{}
	}
	return result
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/enums"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

// articles is a simple in-memory list of articles for inline search.
var articles = []struct {
	Title       string
	Description string
	Content     string
}{
	{"Go", "The Go Programming Language", "Go is an open-source programming language supported by Google."},
	{"Rust", "The Rust Programming Language", "Rust is a systems programming language focused on safety and concurrency."},
	{"Python", "The Python Programming Language", "Python is a high-level, interpreted programming language."},
	{"JavaScript", "The JavaScript Programming Language", "JavaScript is a lightweight, interpreted programming language for the web."},
	{"TypeScript", "TypeScript by Microsoft", "TypeScript is a typed superset of JavaScript that compiles to plain JavaScript."},
	{"Kotlin", "The Kotlin Programming Language", "Kotlin is a modern programming language that targets the JVM and Android."},
	{"Swift", "The Swift Programming Language", "Swift is a powerful and intuitive programming language for Apple platforms."},
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := gramkit.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Handle inline queries (typed — receives *models.InlineQuery directly).
	bot.HandleInlineQueryMatch(gramkit.MatchContains, "", inlineHandler)

	// Handle chosen inline results.
	bot.HandleChosenInlineResult("", func(ctx context.Context, b *gramkit.Bot, cir *models.ChosenInlineResult) {
		log.Printf("User %s chose result: %s", cir.From.FirstName, cir.ResultID)
	})

	// Handle /start in private chat.
	bot.HandleCommand("start", func(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
		b.SendMessage(ctx, params.SendMessage{
			ChatID: msg.Chat.ID,
			Text:   "I'm an inline bot! Type @" + "your_bot_username" + " in any chat to search programming languages.",
		})
	})

	fmt.Println("Inline bot is running...")
	bot.StartPolling(ctx)
}

// inlineHandler responds to inline queries with filtered article results.
func inlineHandler(ctx context.Context, b *gramkit.Bot, iq *models.InlineQuery) {
	query := strings.ToLower(strings.TrimSpace(iq.Query))

	var results []models.InlineQueryResult

	for i, a := range articles {
		if query != "" && !strings.Contains(strings.ToLower(a.Title), query) {
			continue
		}

		results = append(results, models.InlineQueryResultArticle{
			Type:  enums.InlineResultArticle,
			ID:    fmt.Sprintf("%d", i),
			Title: a.Title,
			InputMessageContent: models.InputTextMessageContent{
				MessageText: fmt.Sprintf("*%s*\n\n%s", a.Title, a.Content),
				ParseMode:   "Markdown",
			},
			Description: a.Description,
		})
	}

	b.AnswerInlineQuery(ctx, params.AnswerInlineQuery{
		InlineQueryID: iq.ID,
		Results:       results,
		CacheTime:     10,
		IsPersonal:    true,
	})
}

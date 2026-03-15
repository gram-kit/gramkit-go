package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := gramkit.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Send an inline keyboard when the user types /menu.
	bot.RegisterHandler(gramkit.OnMessage, "/menu", menuHandler)

	// Handle callback queries whose data starts with "pick_".
	bot.RegisterHandlerMatch(gramkit.OnCallbackQuery, gramkit.MatchPrefix, "pick_", pickHandler)

	fmt.Println("Inline keyboard bot is running...")
	bot.StartPolling(ctx)
}

func menuHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Option A", CallbackData: "pick_a"},
				{Text: "Option B", CallbackData: "pick_b"},
			},
			{
				{Text: "Option C", CallbackData: "pick_c"},
			},
		},
	}

	b.SendMessage(ctx, params.SendMessage{
		ChatID:      update.Message.Chat.ID,
		Text:        "Choose an option:",
		ReplyMarkup: keyboard,
	})
}

func pickHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	cb := update.CallbackQuery
	choice := strings.TrimPrefix(cb.Data, "pick_")

	// Acknowledge the callback so the loading indicator disappears.
	b.AnswerCallbackQuery(ctx, params.AnswerCallbackQuery{
		CallbackQueryID: cb.ID,
		Text:            fmt.Sprintf("You picked: %s", choice),
	})

	// Edit the original message to reflect the selection.
	if msg, ok := (*cb.Message).(models.Message); ok {
		b.EditMessageText(ctx, params.EditMessageText{
			ChatID:    msg.Chat.ID,
			MessageID: msg.MessageID,
			Text:      fmt.Sprintf("You selected option %s.", strings.ToUpper(choice)),
		})
	}
}

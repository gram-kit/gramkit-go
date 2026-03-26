# gramkit

A modern Telegram Bot API framework for Go with typed handlers, middleware, hot-reload dev mode, and a CLI tool.

## Features

- Full [Telegram Bot API](https://core.telegram.org/bots/api) coverage
- **Typed handlers** — receive `*Message`, `*CallbackQuery`, `*InlineQuery` instead of raw `*Update`
- **Command routing** — `HandleCommand("start", h)` matches `/start`, `/start@bot`, `/start args`
- Handler-based routing with exact, prefix, and contains matching
- Middleware chain (logger, auth, rate-limit, etc.)
- Long Polling and Webhook support with graceful shutdown
- **Dev mode** — hot-reload with fsnotify, `--debug` (delve), `--watch` (web dashboard)
- `context.Context` in every API call
- Typed enums for API constants
- CLI tool for scaffolding, code generation, and project management

## Installation

```bash
go get github.com/gram-kit/gramkit-go
```

### CLI tool

```bash
go install github.com/gram-kit/gramkit-go/cmd/gramkit@latest
```

## Quick Start

```go
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
	)
	if err != nil {
		log.Fatal(err)
	}

	// Typed command handlers — receive *models.Message directly
	bot.HandleCommand("start", startHandler)
	bot.HandleCommand("help", helpHandler)

	fmt.Println("Bot is running...")
	bot.StartPolling(ctx)
}

func startHandler(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "Hello! I'm a gramkit bot.",
	})
}

func helpHandler(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "Commands:\n/start - Start\n/help - Help",
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
```

## Typed Handlers

Typed handlers receive the specific object instead of the full `*models.Update`:

```go
// Commands — handles /start, /start@bot, /start payload
bot.HandleCommand("start", func(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{ChatID: msg.Chat.ID, Text: "Hello!"})
})

// Messages
bot.HandleMessage("hello", func(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{ChatID: msg.Chat.ID, Text: "Hi!"})
})

// Callback queries
bot.HandleCallbackQuery("btn_yes", func(ctx context.Context, b *gramkit.Bot, cq *models.CallbackQuery) {
	b.AnswerCallbackQuery(ctx, params.AnswerCallbackQuery{CallbackQueryID: cq.ID, Text: "OK"})
})

// Inline queries
bot.HandleInlineQueryMatch(gramkit.MatchContains, "", func(ctx context.Context, b *gramkit.Bot, iq *models.InlineQuery) {
	b.AnswerInlineQuery(ctx, params.AnswerInlineQuery{
		InlineQueryID: iq.ID,
		Results:       results,
	})
})

// With match types
bot.HandleMessageMatch(gramkit.MatchContains, "hello", greetHandler)
bot.HandleCallbackQueryMatch(gramkit.MatchPrefix, "action_", actionHandler)
```

Access the full `*models.Update` from any typed handler via context:

```go
func handler(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	update := gramkit.UpdateFromContext(ctx)
	log.Printf("Update ID: %d", update.UpdateID)
}
```

### All typed handler methods

| Method | Handler receives | Use case |
|--------|-----------------|----------|
| `HandleCommand` | `*Message` | `/start`, `/help`, etc. |
| `HandleMessage` | `*Message` | Text messages |
| `HandleEditedMessage` | `*Message` | Edited messages |
| `HandleChannelPost` | `*Message` | Channel posts |
| `HandleCallbackQuery` | `*CallbackQuery` | Inline keyboard buttons |
| `HandleInlineQuery` | `*InlineQuery` | Inline mode queries |
| `HandleChosenInlineResult` | `*ChosenInlineResult` | Chosen inline results |
| `HandleShippingQuery` | `*ShippingQuery` | Shipping queries |
| `HandlePreCheckoutQuery` | `*PreCheckoutQuery` | Pre-checkout queries |
| `HandlePoll` | `*Poll` | Poll updates |
| `HandlePollAnswer` | `*PollAnswer` | Poll answers |
| `HandleMyChatMember` | `*ChatMemberUpdated` | Bot's member status changes |
| `HandleChatMember` | `*ChatMemberUpdated` | Chat member status changes |
| `HandleChatJoinRequest` | `*ChatJoinRequest` | Join requests |
| `HandleChatBoost` | `*ChatBoostUpdated` | Chat boosts |

## Generic Handler Routing

The low-level API is still available for advanced use cases:

```go
// Exact match
bot.RegisterHandler(gramkit.OnMessage, "/start", startHandler)

// Prefix match — "pick_" catches "pick_a", "pick_b", etc.
bot.RegisterHandlerMatch(gramkit.OnCallbackQuery, gramkit.MatchPrefix, "pick_", pickHandler)

// Contains match
bot.RegisterHandlerMatch(gramkit.OnMessage, gramkit.MatchContains, "hello", greetHandler)

// Wildcard — catch all updates of a type
bot.RegisterHandler(gramkit.OnChatJoinRequest, "*", joinHandler)
```

### Handler types

| Type | Matches |
|------|---------|
| `OnMessage` | `update.Message` |
| `OnCommand` | Bot commands (`/start`, `/help`, etc.) |
| `OnEditedMessage` | `update.EditedMessage` |
| `OnChannelPost` | `update.ChannelPost` |
| `OnCallbackQuery` | `update.CallbackQuery` |
| `OnInlineQuery` | `update.InlineQuery` |
| `OnChatJoinRequest` | `update.ChatJoinRequest` |
| `OnMyChatMember` | `update.MyChatMember` |
| `OnChatMember` | `update.ChatMember` |
| `OnPoll` | `update.Poll` |
| `OnPollAnswer` | `update.PollAnswer` |
| `OnChatBoost` | `update.ChatBoost` |
| `OnMessageReaction` | `update.MessageReaction` |
| `OnBusinessMessage` | `update.BusinessMessage` |
| ...and more | |

### Match types

| Type | Behavior |
|------|----------|
| `MatchExact` | `text == pattern` |
| `MatchPrefix` | `strings.HasPrefix(text, pattern)` |
| `MatchContains` | `strings.Contains(text, pattern)` |

## Middleware

```go
bot, _ := gramkit.New(token,
	gramkit.WithDefaultHandler(handler),
	gramkit.WithMiddleware(logger, auth),
)

// Or add after creation
bot.Use(rateLimiter)
```

Middleware signature:

```go
func logger(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	log.Printf("[%d] %s", update.UpdateID, update.Message.Text)
	next(ctx, b, update) // call next to continue, skip to block
}
```

## Webhook Mode

### Built-in server

```go
bot.StartWebhook(ctx, ":8080")
```

### Custom HTTP server

```go
mux := http.NewServeMux()
mux.Handle("/webhook", bot.WebhookHandler())
mux.HandleFunc("/health", healthCheck)

server := &http.Server{Addr: ":8080", Handler: mux}
server.ListenAndServe()
```

## Options

```go
bot, err := gramkit.New(token,
	gramkit.WithDefaultHandler(handler),     // default handler for unmatched updates
	gramkit.WithMiddleware(mw1, mw2),        // middleware chain
	gramkit.WithHTTPClient(customClient),    // custom *http.Client
	gramkit.WithServerURL("http://localhost:8081"), // local Bot API server
	gramkit.WithPollTimeout(60),             // polling timeout in seconds
)
```

## Enums

Typed constants for Telegram API values:

```go
import "github.com/gram-kit/gramkit-go/enums"

// Send a dice
bot.SendDice(ctx, params.SendDice{
	ChatID: chatID,
	Emoji:  enums.DiceGameDie,
})

// Send a poll
bot.SendPoll(ctx, params.SendPoll{
	ChatID:   chatID,
	Question: "What?",
	Options:  options,
	Type:     enums.PollQuiz,
})

// Check chat member status
if member.Status == enums.MemberAdministrator { ... }

// Allowed updates
bot.StartPolling(ctx, params.GetUpdates{
	AllowedUpdates: []enums.AllowedUpdates{
		enums.UpdateMessage,
		enums.UpdateCallbackQuery,
	},
})
```

## CLI Tool

### Create a new project

```bash
gramkit new my-bot
cd my-bot
# edit .env with your BOT_TOKEN
gramkit run
```

### Dev mode with hot-reload

```bash
gramkit run:dev                    # hot-reload on file changes
gramkit run:dev --watch            # + web dashboard at :4080
gramkit run:dev --debug            # + delve debugger at :2345
gramkit run:dev --debug --watch    # all three
```

### Generate code

```bash
gramkit make:command start         # creates handlers/start.go (typed command handler)
gramkit make:handler profile       # creates handlers/profile.go (typed message handler)
gramkit make:middleware auth        # creates middleware/auth.go
```

### Tools

```bash
gramkit doctor                     # check environment (Go, delve, .env, BOT_TOKEN)
gramkit update                     # self-update CLI to latest version
```

### Webhook management

```bash
gramkit webhook:set https://example.com/webhook
gramkit webhook:info
gramkit webhook:delete
```

## Project Structure

```
gramkit-go/
├── bot.go                  # Bot, Options, StartPolling, StartWebhook
├── router.go               # Handler registration and routing
├── handlers.go             # Typed handler funcs and registration methods
├── context.go              # UpdateFromContext helper
├── middleware.go            # Middleware types and helpers
├── methods.go              # Telegram Bot API methods
├── models/                 # Response types (Update, Message, User, etc.)
├── params/                 # Request parameter types
├── enums/                  # Typed string constants
├── internal/
│   ├── network/            # HTTP client
│   └── devserver/          # Dev watch dashboard (SSE)
├── cmd/gramkit/            # CLI tool
└── examples/               # Example bots
    ├── echo-bot/
    ├── inline-bot/
    ├── inline-keyboard-bot/
    ├── webhook-bot/
    ├── middleware-bot/
    └── webhook-custom-server/
```

## Examples

See the [examples](./examples) directory for complete working bots.

## License

MIT

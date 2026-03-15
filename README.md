# gramkit

A modern Telegram Bot API framework for Go with handler routing, middleware support, and a CLI tool.

## Features

- Full [Telegram Bot API](https://core.telegram.org/bots/api) coverage
- Handler-based routing with pattern matching
- Middleware chain (logger, auth, rate-limit, etc.)
- Long Polling and Webhook support with graceful shutdown
- `context.Context` in every API call
- Functional options for configuration
- Typed enums for API constants
- CLI tool for scaffolding and project management

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
		gramkit.WithDefaultHandler(handler),
	)
	if err != nil {
		log.Fatal(err)
	}

	bot.RegisterHandler(gramkit.OnMessage, "/start", startHandler)

	fmt.Println("Bot is running...")
	bot.StartPolling(ctx)
}

func startHandler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello! I'm a gramkit bot.",
	})
}

func handler(ctx context.Context, b *gramkit.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, params.SendMessage{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
}
```

## Handler Routing

Register handlers for specific update types and patterns:

```go
// Exact match
bot.RegisterHandler(gramkit.OnMessage, "/start", startHandler)
bot.RegisterHandler(gramkit.OnMessage, "/help", helpHandler)

// Prefix match — "pick_" catches "pick_a", "pick_b", etc.
bot.RegisterHandlerMatch(gramkit.OnCallbackQuery, gramkit.MatchPrefix, "pick_", pickHandler)

// Contains match
bot.RegisterHandlerMatch(gramkit.OnMessage, gramkit.MatchContains, "hello", greetHandler)

// Wildcard — catch all updates of a type
bot.RegisterHandler(gramkit.OnChatJoinRequest, "*", joinHandler)
```

### Available handler types

| Type | Matches |
|------|---------|
| `OnMessage` | `update.Message` |
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

### Generate code

```bash
gramkit make:handler profile     # creates handlers/profile.go
gramkit make:middleware auth      # creates middleware/auth.go
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
├── middleware.go            # Middleware types and helpers
├── methods.go              # Telegram Bot API methods
├── models/                 # Response types (Update, Message, User, etc.)
├── params/                 # Request parameter types
├── enums/                  # Typed string constants
├── internal/network/       # HTTP client
├── cmd/gramkit/            # CLI tool
└── examples/               # Example bots
    ├── echo-bot/
    ├── webhook-bot/
    ├── inline-keyboard-bot/
    ├── middleware-bot/
    └── webhook-custom-server/
```

## Examples

See the [examples](./examples) directory for complete working bots.

## License

MIT

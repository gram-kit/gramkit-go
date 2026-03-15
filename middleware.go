package gramkit

import (
	"context"

	"github.com/gram-kit/gramkit-go/models"
)

// MiddlewareFunc is a function that wraps a handler.
// Call next(ctx, b, update) to pass control to the next middleware/handler.
type MiddlewareFunc func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc)

// WithMiddleware adds middleware to the bot during creation.
func WithMiddleware(mw ...MiddlewareFunc) Option {
	return func(b *Bot) {
		b.middlewares = append(b.middlewares, mw...)
	}
}

// Use adds middleware to the bot after creation.
func (b *Bot) Use(mw ...MiddlewareFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.middlewares = append(b.middlewares, mw...)
}

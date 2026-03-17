package gramkit

import (
	"context"

	"github.com/gram-kit/gramkit-go/models"
)

type contextKey struct{}

// UpdateFromContext returns the full *models.Update from the context.
// Useful in typed handlers when you need access to fields beyond the primary object.
func UpdateFromContext(ctx context.Context) *models.Update {
	u, _ := ctx.Value(contextKey{}).(*models.Update)
	return u
}

func withUpdate(ctx context.Context, u *models.Update) context.Context {
	return context.WithValue(ctx, contextKey{}, u)
}

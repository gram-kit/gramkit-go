package params

import (
	"github.com/gram-kit/gramkit-go/enums"
	"github.com/gram-kit/gramkit-go/models"
)

type GetUpdates struct {
	Offset         int64                  `json:"offset,omitempty"`
	Limit          int                    `json:"limit,omitempty"`
	Timeout        int                    `json:"timeout,omitempty"`
	AllowedUpdates []enums.AllowedUpdates `json:"allowed_updates,omitempty"`
}

type SetWebhook struct {
	URL                string                 `json:"url"`
	Certificate        *models.InputFile      `json:"certificate,omitempty"`
	IPAddress          string                 `json:"ip_address,omitempty"`
	MaxConnections     int                    `json:"max_connections,omitempty"`
	AllowedUpdates     []enums.AllowedUpdates `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool                   `json:"drop_pending_updates,omitempty"`
	SecretToken        string                 `json:"secret_token,omitempty"`
}

type DeleteWebhook struct {
	SecretToken string `json:"secret_token,omitempty"`
}

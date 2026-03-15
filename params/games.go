package params

import "github.com/gram-kit/gramkit-go/models"

type SendGame struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id,omitempty"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	GameShortName        string                  `json:"game_short_name"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type SetGameScore struct {
	UserID              int64  `json:"user_id"`
	Scope               int    `json:"scope"`
	Force               bool   `json:"force,omitempty"`
	DisableEditMessages bool   `json:"disable_edit_messages,omitempty"`
	ChatID              any    `json:"chat_id,omitempty"`
	MessageID           int    `json:"message_id,omitempty"`
	InlineMessageID     string `json:"inline_message_id,omitempty"`
}

type GetGameHighScores struct {
	UserID          int64  `json:"user_id"`
	ChatID          any    `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

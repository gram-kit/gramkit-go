package params

import "github.com/gram-kit/gramkit-go/models"

type EditMessageText struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id,omitempty"`
	MessageID            int                          `json:"message_id,omitempty"`
	InlineMessageID      string                       `json:"inline_message_id,omitempty"`
	Text                 string                       `json:"text"`
	ParseMode            string                       `json:"parse_mode,omitempty"`
	Entities             []models.MessageEntity       `json:"entities,omitempty"`
	LinkPreviewOptions   *models.LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCaption struct {
	BusinessConnectionID  string                       `json:"business_connection_id,omitempty"`
	ChatID                any                          `json:"chat_id,omitempty"`
	MessageID             int                          `json:"message_id,omitempty"`
	InlineMessageID       string                       `json:"inline_message_id,omitempty"`
	Caption               string                       `json:"caption,omitempty"`
	ParseMode             string                       `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                         `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageMedia struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id,omitempty"`
	MessageID            int                          `json:"message_id,omitempty"`
	InlineMessageID      string                       `json:"inline_message_id,omitempty"`
	Media                string                       `json:"media,omitempty"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocation struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id,omitempty"`
	MessageID            int                          `json:"message_id,omitempty"`
	InlineMessageID      string                       `json:"inline_message_id,omitempty"`
	Latitude             float64                      `json:"latitude"`
	Longitude            float64                      `json:"longitude"`
	LivePeriod           int                          `json:"live_period,omitempty"`
	HorizontalAccuracy   float64                      `json:"horizontal_accuracy,omitempty"`
	Heading              int                          `json:"heading,omitempty"`
	ProximityAlertRadius int                          `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocation struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id,omitempty"`
	MessageID            int                          `json:"message_id,omitempty"`
	InlineMessageID      string                       `json:"inline_message_id,omitempty"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageChecklist struct {
	BusinessConnectionID string                       `json:"business_connection_id"`
	ChatID               any                          `json:"chat_id"`
	MessageID            int                          `json:"message_id"`
	Checklist            models.InputChecklist        `json:"checklist"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkup struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id,omitempty"`
	MessageID            int                          `json:"message_id,omitempty"`
	InlineMessageID      string                       `json:"inline_message_id,omitempty"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopPoll struct {
	BusinessConnectionID string                       `json:"business_connection_id,omitempty"`
	ChatID               any                          `json:"chat_id"`
	MessageID            int                          `json:"message_id"`
	ReplyMarkup          *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type ApproveSuggestedPost struct {
	ChatID    any   `json:"chat_id"`
	MessageID int   `json:"message_id"`
	StartDate int64 `json:"start_date,omitempty"`
}

type DeclineSuggestedPost struct {
	ChatID    any    `json:"chat_id"`
	MessageID int    `json:"message_id"`
	Comment   string `json:"comment,omitempty"`
}

type DeleteMessage struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

type DeleteMessages struct {
	ChatID     any   `json:"chat_id"`
	MessageIDs []int `json:"message_ids"`
}

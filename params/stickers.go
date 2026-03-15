package params

import (
	"github.com/gram-kit/gramkit-go/enums"
	"github.com/gram-kit/gramkit-go/models"
)

type SendSticker struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  string                          `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                             `json:"direct_messages_topic_id,omitempty"`
	Sticker                 models.InputSticker             `json:"sticker"`
	Emoji                   string                          `json:"emoji,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         models.ReplyParameters          `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type GetStickerSet struct {
	Name string `json:"name"`
}

type GetCustomEmojiStickers struct {
	CustomEmojiIDs []any `json:"custom_emoji_ids"`
}

type UploadStickerFile struct {
	UserID        int64            `json:"user_id"`
	Sticker       models.InputFile `json:"sticker"`
	StickerFormat string           `json:"sticker_format"`
}

type CreateNewStickerSet struct {
	UserID          int64                 `json:"user_id"`
	Name            string                `json:"name"`
	Title           string                `json:"title"`
	Stickers        []models.InputSticker `json:"sticker"`
	StickerType     enums.StickerType     `json:"sticker_type,omitempty"`
	NeedsRepainting bool                  `json:"needs_repainting,omitempty"`
}

type AddStickerToSet struct {
	UserID  int64               `json:"user_id"`
	Name    string              `json:"name"`
	Sticker models.InputSticker `json:"sticker"`
}

type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"`
}

type ReplaceStickerInSet struct {
	UserID     int64  `json:"user_id"`
	Name       string `json:"name"`
	OldSticker string `json:"old_sticker"`
	Sticker    string `json:"sticker"`
}

type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

type SetStickerKeywords struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

type SetStickerMaskPosition struct {
	Sticker      string                `json:"sticker"`
	MaskPosition []models.MaskPosition `json:"mask_position,omitempty"`
}

type SetStickerSetTitle struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type SetStickerSetThumbnail struct {
	Name      string `json:"name"`
	UserID    int64  `json:"user_id"`
	Thumbnail any    `json:"thumbnail,omitempty"`
	Format    string `json:"format"`
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string `json:"name"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type DeleteStickerSet struct {
	Name string `json:"name"`
}

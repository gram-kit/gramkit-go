package models

import "github.com/gram-kit/gramkit-go/enums"

type Sticker struct {
	FileID           string            `json:"file_id"`
	FileUniqueID     string            `json:"file_unique_id"`
	Type             enums.StickerType `json:"type"`
	Width            int               `json:"width"`
	Height           int               `json:"height"`
	IsAnimated       bool              `json:"is_animated"`
	IsVideo          bool              `json:"is_video"`
	Thumbnail        *PhotoSize        `json:"thumbnail,omitempty"`
	Emoji            string            `json:"emoji,omitempty"`
	SetName          string            `json:"set_name,omitempty"`
	PremiumAnimation *File             `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition     `json:"mask_position,omitempty"`
	CustomEmojiID    string            `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool              `json:"needs_repainting,omitempty"`
	FileSize         int               `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	StickerType enums.StickerType `json:"sticker_type"`
	Stickers    []Sticker         `json:"stickers"`
	Thumbnail   *PhotoSize        `json:"thumbnail,omitempty"`
}

type MaskPosition struct {
	Point  enums.MaskPositionPoint `json:"point"`
	XShift float64                 `json:"x_shift"`
	YShift float64                 `json:"y_shift"`
	Scale  float64                 `json:"scale"`
}

type InputSticker struct {
	Sticker      string              `json:"sticker"`
	Format       enums.StickerFormat `json:"format"`
	EmojiList    []string            `json:"emoji_list"`
	MaskPosition *MaskPosition       `json:"mask_position,omitempty"`
	Keywords     []string            `json:"keywords,omitempty"`
}

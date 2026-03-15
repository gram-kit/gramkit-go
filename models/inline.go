package models

import "github.com/gram-kit/gramkit-go/enums"

type InlineQuery struct {
	ID       string         `json:"id"`
	From     User           `json:"from"`
	Query    string         `json:"query"`
	Offset   string         `json:"offset"`
	ChatType enums.ChatType `json:"chat_type,omitempty"`
	Location *Location      `json:"location,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter string      `json:"start_parameter,omitempty"`
}

type InlineQueryResult interface {
	isInlineQueryResult()
}

func (InlineQueryResultCachedAudio) isInlineQueryResult()    {}
func (InlineQueryResultCachedDocument) isInlineQueryResult() {}
func (InlineQueryResultCachedGif) isInlineQueryResult()      {}
func (InlineQueryResultCachedMpeg4Gif) isInlineQueryResult() {}
func (InlineQueryResultCachedPhoto) isInlineQueryResult()    {}
func (InlineQueryResultCachedSticker) isInlineQueryResult()  {}
func (InlineQueryResultCachedVideo) isInlineQueryResult()    {}
func (InlineQueryResultCachedVoice) isInlineQueryResult()    {}
func (InlineQueryResultArticle) isInlineQueryResult()        {}
func (InlineQueryResultAudio) isInlineQueryResult()          {}
func (InlineQueryResultContact) isInlineQueryResult()        {}
func (InlineQueryResultGame) isInlineQueryResult()           {}
func (InlineQueryResultDocument) isInlineQueryResult()       {}
func (InlineQueryResultGif) isInlineQueryResult()            {}
func (InlineQueryResultLocation) isInlineQueryResult()       {}
func (InlineQueryResultMpeg4Gif) isInlineQueryResult()       {}
func (InlineQueryResultPhoto) isInlineQueryResult()          {}
func (InlineQueryResultVenue) isInlineQueryResult()          {}
func (InlineQueryResultVideo) isInlineQueryResult()          {}
func (InlineQueryResultVoice) isInlineQueryResult()          {}

type InlineQueryResultArticle struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	Title               string                      `json:"title"`
	InputMessageContent InputMessageContent         `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	URL                 string                      `json:"url,omitempty"`
	Description         string                      `json:"description,omitempty"`
	ThumbnailURL        string                      `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                         `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                         `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultPhoto struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	PhotoURL              string                      `json:"photo_url"`
	ThumbnailURL          string                      `json:"thumbnail_url"`
	PhotoWidth            int                         `json:"photo_width,omitempty"`
	PhotoHeight           int                         `json:"photo_height,omitempty"`
	Title                 string                      `json:"title,omitempty"`
	Description           string                      `json:"description,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent         `json:"input_message_content"`
}

type InlineQueryResultGif struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	GifURL                string                      `json:"gif_url"`
	GifWidth              int                         `json:"gif_width,omitempty"`
	GifHeight             int                         `json:"gif_height,omitempty"`
	GifDuration           int                         `json:"gif_duration,omitempty"`
	ThumbnailURL          string                      `json:"thumbnail_url,omitempty"`
	ThumbnailMimeType     string                      `json:"thumbnail_mime_type,omitempty"`
	Title                 string                      `json:"title,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	Mpeg4URL              string                      `json:"mpeg4_url"`
	Mpeg4Width            int                         `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int                         `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         int                         `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string                      `json:"thumbnail_url"`
	ThumbnailMimeType     string                      `json:"thumbnail_mime_type,omitempty"`
	Title                 string                      `json:"title,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	VideoURL              string                      `json:"video_url"`
	MimeType              string                      `json:"mime_type"`
	ThumbnailURL          string                      `json:"thumbnail_url"`
	Title                 string                      `json:"title"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	VideoWidth            int                         `json:"video_width,omitempty"`
	VideoHeight           int                         `json:"video_height,omitempty"`
	VideoDuration         int                         `json:"video_duration,omitempty"`
	Description           string                      `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	AudioURL            string                      `json:"audio_url"`
	Title               string                      `json:"title"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	Performer           string                      `json:"performer,omitempty"`
	AudioDuration       int                         `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	VoiceURL            string                      `json:"voice_url"`
	Title               string                      `json:"title"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	VoiceDuration       int                         `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	Title               string                      `json:"title"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	DocumentURL         string                      `json:"document_url"`
	MimeType            string                      `json:"mime_type"`
	Description         string                      `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
	ThumbnailURL        string                      `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                         `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                         `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title,omitempty"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                   `json:"live_period,omitempty"`
	Heading              int                   `json:"heading,omitempty"`
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	Latitude            float64                     `json:"latitude"`
	Longitude           float64                     `json:"longitude"`
	Title               string                      `json:"title,omitempty"`
	Address             string                      `json:"address"`
	FoursquareID        string                      `json:"foursquare_id,omitempty"`
	FoursquareType      string                      `json:"foursquare_type,omitempty"`
	GooglePlaceID       string                      `json:"google_place_id,omitempty"`
	GooglePlaceType     string                      `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
	ThumbnailURL        string                      `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                         `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                         `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	PhoneNumber         string                      `json:"phone_number"`
	FirstName           string                      `json:"first_name"`
	LastName            string                      `json:"last_name,omitempty"`
	VCard               string                      `json:"v_card,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
	ThumbnailURL        string                      `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                         `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                         `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineQueryResultCachedPhoto struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	PhotoFileID           string                      `json:"photo_file_id"`
	Title                 string                      `json:"title,omitempty"`
	Description           string                      `json:"description,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	GifFileId             string                      `json:"gif_file_id"`
	Title                 string                      `json:"title,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	Mpeg4FileID           string                      `json:"mpeg4_file_id"`
	Title                 string                      `json:"title,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	StickerFileID       string                      `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	Title               string                      `json:"title"`
	DocumentFileID      string                      `json:"document_file_id"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	MimeType            string                      `json:"mime_type"`
	Description         string                      `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                  enums.InlineQueryResultType `json:"type"`
	ID                    string                      `json:"id"`
	VideoFileID           string                      `json:"video_file_id"`
	Title                 string                      `json:"title"`
	Description           string                      `json:"description,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             string                      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity             `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	VoiceFileID         string                      `json:"voice_file_id"`
	Title               string                      `json:"title"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                enums.InlineQueryResultType `json:"type"`
	ID                  string                      `json:"id"`
	AudioFileID         string                      `json:"audio_file_id"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity             `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}

type InputMessageContent interface {
	isInputMessageContent()
}

func (InputTextMessageContent) isInputMessageContent()     {}
func (InputLocationMessageContent) isInputMessageContent() {}
func (InputVenueMessageContent) isInputMessageContent()    {}
func (InputContactMessageContent) isInputMessageContent()  {}
func (InputInvoiceMessageContent) isInputMessageContent()  {}

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int64  `json:"expiration_date"`
}

package params

import (
	"github.com/gram-kit/gramkit-go/enums"
	"github.com/gram-kit/gramkit-go/models"
)

type SendMessage struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Text                    string                          `json:"text"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	Entities                []models.MessageEntity          `json:"entities,omitempty"`
	LinkPreviewOptions      *models.LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type ForwardMessage struct {
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	FromChatID              any                             `json:"from_chat_id"`
	VideoStartTimestamp     int64                           `json:"video_start_timestamp,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	MessageID               int                             `json:"message_id"`
}

type ForwardMessages struct {
	ChatID               any   `json:"chat_id"`
	MessageThreadID      int   `json:"message_thread_id,omitempty"`
	DirectMessageTopicID int   `json:"direct_message_topic_id,omitempty"`
	FromChatID           any   `json:"from_chat_id"`
	MessageIDs           []int `json:"message_ids"`
	DisableNotification  bool  `json:"disable_notification,omitempty"`
	ProtectContent       bool  `json:"protect_content,omitempty"`
}

type CopyMessage struct {
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	FromChatID              any                             `json:"from_chat_id"`
	MessageID               int                             `json:"message_id"`
	VideoStartTimestamp     int64                           `json:"video_start_timestamp,omitempty"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                            `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type CopyMessages struct {
	ChatID               any   `json:"chat_id"`
	MessageThreadID      int   `json:"message_thread_id,omitempty"`
	DirectMessageTopicID int   `json:"direct_message_topic_id,omitempty"`
	FromChatID           any   `json:"from_chat_id"`
	MessageIDs           []int `json:"message_ids"`
	DisableNotification  bool  `json:"disable_notification,omitempty"`
	ProtectContent       bool  `json:"protect_content,omitempty"`
	RemoveCaption        bool  `json:"remove_caption,omitempty"`
}

type SendPhoto struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Photo                   any                             `json:"photo"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                            `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                            `json:"has_spoiler,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendAudio struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Audio                   any                             `json:"audio"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int                             `json:"duration,omitempty"`
	Performer               string                          `json:"performer,omitempty"`
	Title                   string                          `json:"title,omitempty"`
	Thumbnail               any                             `json:"thumbnail,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendDocument struct {
	BusinessConnectionID        string                          `json:"business_connection_id,omitempty"`
	ChatID                      any                             `json:"chat_id"`
	MessageThreadID             int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID        int                             `json:"direct_message_topic_id,omitempty"`
	Document                    any                             `json:"document"`
	Thumbnail                   any                             `json:"thumbnail,omitempty"`
	Caption                     string                          `json:"caption,omitempty"`
	ParseMode                   string                          `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity          `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                            `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                            `json:"disable_notification,omitempty"`
	ProtectContent              bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast          bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID             string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters     *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters             *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup                 models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendVideo struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Video                   any                             `json:"video"`
	Duration                int                             `json:"duration,omitempty"`
	Width                   int                             `json:"width,omitempty"`
	Height                  int                             `json:"height,omitempty"`
	Thumbnail               any                             `json:"thumbnail,omitempty"`
	Cover                   any                             `json:"cover,omitempty"`
	StartTimestamp          int                             `json:"start_timestamp,omitempty"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                            `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                            `json:"has_spoiler,omitempty"`
	SupportStreaming        bool                            `json:"support_streaming,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendAnimation struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Animation               any                             `json:"animation"`
	Duration                int                             `json:"duration,omitempty"`
	Width                   int                             `json:"width,omitempty"`
	Height                  int                             `json:"height,omitempty"`
	Thumbnail               any                             `json:"thumbnail,omitempty"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                            `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                            `json:"has_spoiler,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendVoice struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Voice                   any                             `json:"voice"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int                             `json:"duration,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendVideoNote struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	VideoNote               any                             `json:"video_note"`
	Duration                int                             `json:"duration,omitempty"`
	Length                  int                             `json:"length,omitempty"`
	Thumbnail               any                             `json:"thumbnail,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendPaidMedia struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	StarCount               int64                           `json:"star_count"`
	Media                   []models.InputPaidMedia         `json:"media"`
	Payload                 string                          `json:"payload,omitempty"`
	Caption                 string                          `json:"caption,omitempty"`
	ParseMode               string                          `json:"parse_mode,omitempty"`
	CaptionEntities         []models.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                            `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendMediaGroup struct {
	BusinessConnectionID string                   `json:"business_connection_id,omitempty"`
	ChatID               any                      `json:"chat_id"`
	MessageThreadID      int                      `json:"message_thread_id,omitempty"`
	DirectMessageTopicID int                      `json:"direct_message_topic_id,omitempty"`
	Media                []models.InputMediaGroup `json:"media"`
	DisableNotification  bool                     `json:"disable_notification,omitempty"`
	ProtectContent       bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                   `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters  `json:"reply_parameters,omitempty"`
}

type SendLocation struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Latitude                float64                         `json:"latitude"`
	Longitude               float64                         `json:"longitude"`
	HorizontalAccuracy      float64                         `json:"horizontal_accuracy"`
	LivePeriod              int                             `json:"live_period,omitempty"`
	Heading                 int                             `json:"heading,omitempty"`
	ProximityAlertRadius    int                             `json:"proximity_alert_radius,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendVenue struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Latitude                float64                         `json:"latitude"`
	Longitude               float64                         `json:"longitude"`
	Title                   string                          `json:"title,omitempty"`
	Address                 string                          `json:"address,omitempty"`
	FoursquareID            string                          `json:"foursquare_id,omitempty"`
	FoursquareType          string                          `json:"foursquare_type,omitempty"`
	GooglePlaceID           string                          `json:"google_place_id,omitempty"`
	GooglePlaceType         string                          `json:"google_place_type,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendContact struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	PhoneNumber             string                          `json:"phone_number"`
	FirstName               string                          `json:"first_name"`
	LastName                string                          `json:"last_name,omitempty"`
	VCard                   string                          `json:"vcard,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendPoll struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	Question                string                          `json:"question"`
	QuestionParseMode       string                          `json:"question_parse_mode,omitempty"`
	QuestionEntities        []models.MessageEntity          `json:"question_entities,omitempty"`
	Options                 []models.InputPollOption        `json:"options"`
	IsAnonymous             bool                            `json:"is_anonymous,omitempty"`
	Type                    enums.PollType                  `json:"type,omitempty"`
	AllowMultipleAnswers    bool                            `json:"allow_multiple_answers,omitempty"`
	CorrectOptionID         int                             `json:"correct_option_id,omitempty"`
	ExplanationParseMode    string                          `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities     []models.MessageEntity          `json:"explanation_entities,omitempty"`
	OpenPeriod              int                             `json:"open_period,omitempty"`
	CloseDate               int64                           `json:"close_date,omitempty"`
	IsClosed                bool                            `json:"is_closed,omitempty"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendChecklist struct {
	BusinessConnectionID string                  `json:"business_connection_id"`
	ChatID               any                     `json:"chat_id"`
	Checklist            models.InputChecklist   `json:"checklist"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type SendDice struct {
	BusinessConnectionID    string                          `json:"business_connection_id,omitempty"`
	ChatID                  any                             `json:"chat_id"`
	MessageThreadID         int                             `json:"message_thread_id,omitempty"`
	DirectMessageTopicID    int                             `json:"direct_message_topic_id,omitempty"`
	Emoji                   enums.DiceValue                 `json:"emoji"`
	DisableNotification     bool                            `json:"disable_notification,omitempty"`
	ProtectContent          bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             models.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendMessageDraft struct {
	ChatID          any                    `json:"chat_id"`
	MessageThreadID int                    `json:"message_thread_id,omitempty"`
	DraftID         int64                  `json:"draft_id"`
	Text            string                 `json:"text"`
	ParseMode       string                 `json:"parse_mode,omitempty"`
	Entities        []models.MessageEntity `json:"entities,omitempty"`
}

type SendChatAction struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageThreadID      int    `json:"message_thread_id,omitempty"`
	Action               string `json:"action"`
}

type SetMessageReaction struct {
	ChatID    any                   `json:"chat_id"`
	MessageID int                   `json:"message_id"`
	Reaction  []models.ReactionType `json:"reaction,omitempty"`
	IsBig     bool                  `json:"is_big"`
}

type GetUserProfilePhotos struct {
	UserID int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

type GetUserProfileAudios struct {
	UserID int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

type SetUserEmojiStatus struct {
	UserID                    int64  `json:"user_id"`
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id"`
	EmojiStatusExpirationDate string `json:"emoji_status_expiration_date"`
}

type GetFile struct {
	FileID string `json:"file_id"`
}

type BanChatMember struct {
	ChatID         any   `json:"chat_id"`
	UserID         int64 `json:"user_id"`
	UntilDate      int64 `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

type UnbanChatMember struct {
	ChatID       any   `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

type RestrictChatMember struct {
	ChatID                        any                    `json:"chat_id"`
	UserID                        int64                  `json:"user_id"`
	Permissions                   models.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                   `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64                  `json:"until_date,omitempty"`
}

type PromoteChatMember struct {
	ChatID                  any   `json:"chat_id"`
	UserID                  int64 `json:"user_id"`
	IsAnonymous             bool  `json:"is_anonymous,omitempty"`
	CanManageChat           bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo           bool  `json:"can_change_info,omitempty"`
	CanInviteUsers          bool  `json:"can_invite_users,omitempty"`
	CanPostStories          bool  `json:"can_post_stories,omitempty"`
	CanEditStories          bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool  `json:"can_post_messages,omitempty"`
	CanEditMessages         bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool  `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool  `json:"can_manage_direct_messages,omitempty"`
	CanManageTags           bool  `json:"can_manage_tags,omitempty"`
}

type SetChatAdministratorCustomTitle struct {
	ChatID      any    `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

type SetChatMemberTag struct {
	ChatID any    `json:"chat_id"`
	UserID int64  `json:"user_id"`
	Tag    string `json:"tag,omitempty"`
}

type BanChatSenderChat struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

type UnbanChatSenderChat struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

type SetChatPermissions struct {
	ChatID                        any                    `json:"chat_id"`
	Permissions                   models.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                   `json:"use_independent_chat_permissions,omitempty"`
}

type ExportChatInviteLink struct {
	ChatID any `json:"chat_id"`
}

type CreateChatInviteLink struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type EditChatInviteLink struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type CreateChatSubscriptionInviteLink struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod string `json:"subscription_period"`
	SubscriptionURL    string `json:"subscription_url"`
}

type EditChatSubscriptionInviteLink struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
	Name       string `json:"name,omitempty"`
}

type RevokeChatInviteLink struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

type ApproveChatJoinRequest struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type DeclineChatJoinRequest struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type SetChatPhoto struct {
	ChatID any              `json:"chat_id"`
	Photo  models.InputFile `json:"photo"`
}

type DeleteChatPhoto struct {
	ChatID any `json:"chat_id"`
}

type SetChatTitle struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

type SetChatDescription struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"description"`
}

type PinChatMessage struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int    `json:"message_id"`
	DisableNotification  bool   `json:"disable_notification,omitempty"`
}

type UnpinChatMessage struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int    `json:"message_id,omitempty"`
}
type UnpinAllChatMessage struct {
	ChatID any `json:"chat_id"`
}

type LeaveChat struct {
	ChatID any `json:"chat_id"`
}

type GetChat struct {
	ChatID any `json:"chat_id"`
}

type GetChatAdministrators struct {
	ChatID any `json:"chat_id"`
}

type GetChatMemberCount struct {
	ChatID any `json:"chat_id"`
}

type GetChatMember struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type SetChatStickerSet struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type DeleteChatStickerSet struct {
	ChatID any `json:"chat_id"`
}

type CreateForumTopic struct {
	ChatID            any    `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type EditForumTopic struct {
	ChatID            any    `json:"chat_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type CloseForumTopic struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type ReopenForumTopic struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type DeleteForumTopic struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type UnpinAllForumTopicMessages struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type EditGeneralForumTopic struct {
	ChatID any    `json:"chat_id"`
	Name   string `json:"name,omitempty"`
}

type CloseGeneralForumTopic struct {
	ChatID any `json:"chat_id"`
}

type ReopenGeneralForumTopic struct {
	ChatID any `json:"chat_id"`
}

type HideGeneralForumTopic struct {
	ChatID any `json:"chat_id"`
}

type UnhideGeneralForumTopic struct {
	ChatID any `json:"chat_id"`
}

type UnpinAllGeneralForumTopicMessages struct {
	ChatID any `json:"chat_id"`
}

type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

type GetUserChatBoosts struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type GetBusinessConnection struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

type SetMyCommands struct {
	Commands     []models.BotCommand    `json:"commands"`
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type DeleteMyCommands struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type GetMyCommands struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type SetMyName struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyDescription struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyProfilePhoto struct {
	Photo models.InputProfilePhoto `json:"photo"`
}

type SetChatMenuButton struct {
	ChatID     any               `json:"chat_id,omitempty"`
	MenuButton models.MenuButton `json:"menu_button,omitempty"`
}

type GetChatMenuButton struct {
	ChatID any `json:"chat_id,omitempty"`
}

type SetMyDefaultAdministratorRights struct {
	Rights      models.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                           `json:"for_channels,omitempty"`
}

type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

type SendGift struct {
	UserID        int64                  `json:"user_id,omitempty"`
	ChatID        any                    `json:"chat_id,omitempty"`
	GiftID        string                 `json:"gift_id"`
	PayForUpgrade bool                   `json:"pay_for_upgrade,omitempty"`
	Text          string                 `json:"text,omitempty"`
	TextParseMode string                 `json:"text_parse_mode,omitempty"`
	TextEntities  []models.MessageEntity `json:"text_entities,omitempty"`
}

type GiftPremiumSubscription struct {
	UserID        int64                  `json:"user_id"`
	MonthCount    int                    `json:"month_count"`
	StarCount     int64                  `json:"star_count"`
	Text          string                 `json:"text,omitempty"`
	TextParseMode string                 `json:"text_parse_mode,omitempty"`
	TextEntities  []models.MessageEntity `json:"text_entities,omitempty"`
}

type VerifyUser struct {
	UserID            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

type VerifyChat struct {
	UserID            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

type RemoveUserVerification struct {
	UserID int64 `json:"user_id"`
}

type RemoveChatVerification struct {
	ChatID any `json:"chat_id"`
}

type ReadBusinessMessage struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ChatID               any    `json:"chat_id"`
	MessageID            int    `json:"message_id"`
}

type DeleteBusinessMessages struct {
	BusinessConnectionID string `json:"business_connection_id"`
	MessageIDs           []int  `json:"message_ids"`
}

type SetBusinessAccountName struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name,omitempty"`
}

type SetBusinessAccountUsername struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Username             string `json:"username,omitempty"`
}

type SetBusinessAccountBio struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Bio                  string `json:"bio,omitempty"`
}

type SetBusinessAccountProfilePhoto struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	Photo                models.InputProfilePhoto `json:"photo"`
	IsPublic             bool                     `json:"is_public,omitempty"`
}

type RemoveBusinessAccountProfilePhoto struct {
	BusinessConnectionID string `json:"business_connection_id"`
	IsPublic             bool   `json:"is_public,omitempty"`
}

type SetBusinessAccountGiftSettings struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	ShowGiftButton       bool                     `json:"show_gift_button"`
	AcceptedGiftTypes    models.AcceptedGiftTypes `json:"accepted_gift_types"`
}

type GetBusinessAccountStarBalance struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

type TransferBusinessAccountStars struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StarCount            int64  `json:"star_count"`
}

type GetBusinessAccountGifts struct {
	BusinessConnectionID        string `json:"business_connection_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

type GetUserGifts struct {
	UserID                      int64  `json:"user_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

type GetChatGifts struct {
	ChatID                      any    `json:"chat_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

type ConvertGiftToStars struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
}

type UpgradeGift struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	KeepOriginalDetails  bool   `json:"keep_original_details,omitempty"`
	StarCount            int64  `json:"star_count,omitempty"`
}

type TransferGift struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	NewOwnerChatID       string `json:"new_owner_chat_id"`
	StarCount            int64  `json:"star_count,omitempty"`
}

type PostStory struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	Content              models.InputStoryContent `json:"content"`
	ActivePeriod         int64                    `json:"active_period"`
	Caption              string                   `json:"caption,omitempty"`
	ParseMode            string                   `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []models.StoryArea       `json:"areas,omitempty"`
	PostToChatPage       bool                     `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool                     `json:"protect_content,omitempty"`
}

type RepostStory struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FromChatID           int64  `json:"from_chat_id"`
	FromStoryID          int64  `json:"from_story_id"`
	ActivePeriod         int64  `json:"active_period"`
	PostToChatPage       bool   `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool   `json:"protect_content,omitempty"`
}

type EditStory struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	StoryID              int64                    `json:"story_id"`
	Content              models.InputStoryContent `json:"content"`
	Caption              string                   `json:"caption,omitempty"`
	ParseMode            string                   `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []models.StoryArea       `json:"areas,omitempty"`
}

type DeleteStory struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StoryID              int64  `json:"story_id"`
}

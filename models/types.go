package models

import (
	"encoding/json"

	"github.com/gram-kit/gramkit-go/enums"
)

type User struct {
	ID                        int64  `json:"id"`
	IsBot                     bool   `json:"is_bot"`
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name,omitempty"`
	Username                  string `json:"username,omitempty"`
	LanguageCode              string `json:"language_code,omitempty"`
	IsPremium                 bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu     bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups             bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages   bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries     bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness      bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp             bool   `json:"has_main_web_app,omitempty"`
	HasTopicsEnabled          bool   `json:"has_topics_enabled,omitempty"`
	AllowsUsersToCreateTopics bool   `json:"allows_users_to_create_topics,omitempty"`
}

type Chat struct {
	ID               int64          `json:"id"`
	Type             enums.ChatType `json:"type"`
	Title            string         `json:"title,omitempty"`
	Username         string         `json:"username,omitempty"`
	FirstName        string         `json:"first_name,omitempty"`
	LastName         string         `json:"last_name,omitempty"`
	IsForum          bool           `json:"is_forum,omitempty"`
	IsDirectMessages bool           `json:"is_direct_messages,omitempty"`
}

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               enums.ChatType        `json:"type"`
	Title                              string                `json:"title,omitempty"`
	Username                           string                `json:"username,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	IsDirectMessages                   bool                  `json:"is_direct_messages,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	Description                        string                `json:"description,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	Rating                             *UserRating           `json:"rating,omitempty"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio,omitempty"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors,omitempty"`
	PaidMessageStarCount               int64                 `json:"paid_message_star_count,omitempty"`
}

type Message struct {
	MessageID                     int                            `json:"message_id"`
	MessageThreadID               int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	SenderTag                     string                         `json:"sender_tag,omitempty"`
	Date                          int64                          `json:"date"`
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`
	Chat                          Chat                           `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"-"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ReplyToChecklistTaskID        int64                          `json:"reply_to_checklist_task_id,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int64                          `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	IsPaidPost                    bool                           `json:"is_paid_post,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	PaidStarCount                 int64                          `json:"paid_star_count,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	EffectID                      string                         `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"` // Note: Sticker struct not defined here
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	Checklist                     *Checklist                     `json:"checklist,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left,omitempty"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"-"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`        // Note: GiftInfo struct not defined here
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"` // Note: UniqueGiftInfo struct not defined here
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done,omitempty"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added,omitempty"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved,omitempty"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed,omitempty"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined,omitempty"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid,omitempty"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"` // Note: InlineKeyboardMarkup struct not defined here
}

type MessageID struct {
	MessageID int `json:"message_id"`
}

type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	MessageID int   `json:"message_id"`
	Date      int64 `json:"date"` // Всегда 0
}

type MaybeInaccessibleMessage interface {
	isMaybeInaccessibleMessage()
}

func (Message) isMaybeInaccessibleMessage()             {}
func (InaccessibleMessage) isMaybeInaccessibleMessage() {}

type MessageEntity struct {
	Type           enums.MessageEntityType `json:"type"`
	Offset         int                     `json:"offset"`
	Length         int                     `json:"length"`
	URL            string                  `json:"url,omitempty"`
	User           *User                   `json:"user,omitempty"`
	Language       string                  `json:"language,omitempty"`
	CustomEmojiID  string                  `json:"custom_emoji_id,omitempty"`
	UnixTime       int64                   `json:"unix_time,omitempty"`
	DateTimeFormat string                  `json:"date_time_format,omitempty"`
}

type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Position int             `json:"position"`
	IsManual bool            `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"-"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageID          int                 `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Checklist          *Checklist          `json:"checklist,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

type ReplyParameters struct {
	MessageID                int             `json:"message_id"`
	ChatID                   interface{}     `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    string          `json:"quote,omitempty"`
	QuoteParseMode           string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int             `json:"quote_position,omitempty"`
	ChecklistTaskID          int64           `json:"checklist_task_id,omitempty"`
}

type MessageOrigin interface {
	isMessageOrigin()
}

func (MessageOriginUser) isMessageOrigin()       {}
func (MessageOriginHiddenUser) isMessageOrigin() {}
func (MessageOriginChannel) isMessageOrigin()    {}
func (MessageOriginChat) isMessageOrigin()       {}

type MessageOriginUser struct {
	Type       enums.MessageOriginType `json:"type"`
	Date       int64                   `json:"date"`
	SenderUser User                    `json:"sender_user"`
}

type MessageOriginHiddenUser struct {
	Type           enums.MessageOriginType `json:"type"`
	Date           int64                   `json:"date"`
	SenderUserName string                  `json:"sender_user_name"`
}

type MessageOriginChat struct {
	Type            enums.MessageOriginType `json:"type"`
	Date            int64                   `json:"date"`
	SenderChat      Chat                    `json:"sender_chat"`
	AuthorSignature string                  `json:"author_signature,omitempty"`
}

type MessageOriginChannel struct {
	Type            enums.MessageOriginType `json:"type"`
	Date            int64                   `json:"date"`
	Chat            Chat                    `json:"chat"`
	MessageID       int                     `json:"message_id"`
	AuthorSignature string                  `json:"author_signature,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Story struct {
	Chat Chat  `json:"chat"`
	ID   int64 `json:"id"`
}

type VideoQuality struct {
	FileId       string                  `json:"file_id"`
	FileUniqueId string                  `json:"file_unique_id"`
	Width        int                     `json:"width"`
	Height       int                     `json:"height"`
	Codec        enums.VideoQualityCodec `json:"codec"`
	FileSize     int                     `json:"file_size"`
}

type Video struct {
	FileID         string         `json:"file_id"`
	FileUniqueID   string         `json:"file_unique_id"`
	Width          int            `json:"width"`
	Height         int            `json:"height"`
	Duration       int            `json:"duration"`
	Thumbnail      *PhotoSize     `json:"thumbnail,omitempty"`
	Cover          []PhotoSize    `json:"cover,omitempty"`
	StartTimestamp int64          `json:"start_timestamp,omitempty"`
	Qualities      []VideoQuality `json:"qualities,omitempty"`
	FileName       string         `json:"file_name,omitempty"`
	MimeType       string         `json:"mime_type,omitempty"`
	FileSize       int            `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"` // Длина стороны квадрата в пикселях
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"`
	PaidMedia []PaidMedia `json:"-"`
}

type PaidMedia interface {
	isPaidMedia()
}

func (PaidMediaPreview) isPaidMedia() {}
func (PaidMediaPhoto) isPaidMedia()   {}
func (PaidMediaVideo) isPaidMedia()   {}

type PaidMediaPreview struct {
	Type     enums.PaidMediaType `json:"type"`
	Width    int                 `json:"width,omitempty"`
	Height   int                 `json:"height,omitempty"`
	Duration int                 `json:"duration,omitempty"`
}

type PaidMediaPhoto struct {
	Type  enums.PaidMediaType `json:"type"`
	Photo []PhotoSize         `json:"photo"`
}

type PaidMediaVideo struct {
	Type  enums.PaidMediaType `json:"type"`
	Video Video               `json:"video"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji enums.DiceValue `json:"emoji"`
	Value int             `json:"value"`
}

type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`
}

type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  enums.PollType  `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int64           `json:"close_date,omitempty"`
}

type ChecklistTask struct {
	ID              int64           `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User           `json:"completed_by_user,omitempty"`
	CompletedByChat *Chat           `json:"completed_by_chat,omitempty"`
	CompletionDate  int64           `json:"completion_date,omitempty"`
}

type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done,omitempty"`
}

type InputChecklistTask struct {
	ID           int64           `json:"id"`
	Text         string          `json:"text"`
	ParseMode    string          `json:"parse_mode,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
}

type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIDs    []int64  `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIDs []int64  `json:"marked_as_not_done_task_ids,omitempty"`
}

type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        Location                  `json:"location"`
	Title           string                    `json:"title"`
	Address         string                    `json:"address"`
	FoursquareID    string                    `json:"foursquare_id,omitempty"`
	FoursquareType  enums.VenueFoursquareType `json:"foursquare_type,omitempty"`
	GooglePlaceID   string                    `json:"google_place_id,omitempty"`
	GooglePlaceType string                    `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type BackgroundFill interface {
	isBackgroundFill()
}

func (BackgroundFillSolid) isBackgroundFill()            {}
func (BackgroundFillGradient) isBackgroundFill()         {}
func (BackgroundFillFreeformGradient) isBackgroundFill() {}

type BackgroundFillSolid struct {
	Type  enums.BackgroundFillType `json:"type"`
	Color int                      `json:"color"`
}

type BackgroundFillGradient struct {
	Type          enums.BackgroundFillType `json:"type"`
	TopColor      int                      `json:"top_color"`
	BottomColor   int                      `json:"bottom_color"`
	RotationAngle int                      `json:"rotation_angle"`
}

type BackgroundFillFreeformGradient struct {
	Type   enums.BackgroundFillType `json:"type"`
	Colors []int                    `json:"colors"`
}

type BackgroundType interface {
	isBackgroundType()
}

func (BackgroundTypeFill) isBackgroundType()      {}
func (BackgroundTypeWallpaper) isBackgroundType() {}
func (BackgroundTypePattern) isBackgroundType()   {}
func (BackgroundTypeChatTheme) isBackgroundType() {}

type BackgroundTypeFill struct {
	Type             enums.BackgroundType `json:"type"`
	Fill             BackgroundFill       `json:"-"`
	DarkThemeDimming int                  `json:"dark_theme_dimming"`
}

type BackgroundTypeWallpaper struct {
	Type             enums.BackgroundType `json:"type"`
	Document         Document             `json:"document"`
	DarkThemeDimming int                  `json:"dark_theme_dimming"`
	IsBlurred        bool                 `json:"is_blurred,omitempty"`
	IsMoving         bool                 `json:"is_moving,omitempty"`
}

type BackgroundTypePattern struct {
	Type       enums.BackgroundType `json:"type"`
	Document   Document             `json:"document"`
	Fill       BackgroundFill       `json:"-"`
	Intensity  int                  `json:"intensity"`
	IsInverted bool                 `json:"is_inverted,omitempty"`
	IsMoving   bool                 `json:"is_moving,omitempty"`
}

type BackgroundTypeChatTheme struct {
	Type      enums.BackgroundType `json:"type"`
	ThemeName string               `json:"theme_name"`
}

type ChatBackground struct {
	Type BackgroundType `json:"-"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

type ForumTopicClosed struct{}

type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}

type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName string      `json:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestID int          `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

type ChatShared struct {
	RequestID int         `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     string      `json:"title,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

type VideoChatStarted struct{}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type PaidMessagePriceChanged struct {
	PaidMessageStarCount int64 `json:"paid_message_star_count"`
}

type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool  `json:"are_direct_messages_enabled,omitempty"`
	DirectMessageStarCount   int64 `json:"direct_message_star_count"`
}

type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message"`
	Price                *SuggestedPostPrice `json:"price"`
	SendDate             int64               `json:"send_date"`
}

type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message"`
	Price                *SuggestedPostPrice `json:"price"`
}

type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Comment              string   `json:"comment,omitempty"`
}

type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message"`
	Currency             string      `json:"currency,omitempty"`
	Amount               int         `json:"amount"`
	StarAmount           *StarAmount `json:"star_amount"` // StarAmount - сложный тип
}

type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Reason               string   `json:"reason,omitempty"`
}

type GiveawayCreated struct {
	PrizeStarCount int64 `json:"prize_star_count,omitempty"`
}

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                int64    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          Chat   `json:"chat"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int64  `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int64  `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
	PrizeDescription              string `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count"`
	GiveawayMessage     *Message `json:"giveaway_message"`
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	URL              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date"`
}

type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date"`
}

type DirectMessagesTopic struct {
	TopicID int64 `json:"topic_id"`
	User    User  `json:"user"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type UserProfileAudios struct {
	TotalCount int     `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type ReplyMarkup interface {
	isReplyMarkup()
}

func (InlineKeyboardMarkup) isReplyMarkup() {}
func (ReplyKeyboardMarkup) isReplyMarkup()  {}
func (ReplyKeyboardRemove) isReplyMarkup()  {}
func (ForceReply) isReplyMarkup()           {}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text              string                      `json:"text"`
	IconCustomEmojiID string                      `json:"icon_custom_emoji_id,omitempty"`
	Style             enums.KeyboardButtonStyle   `json:"style,omitempty"`
	RequestUsers      *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat       *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact    bool                        `json:"request_contact,omitempty"`
	RequestLocation   bool                        `json:"request_location,omitempty"`
	RequestPoll       *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp            *WebAppInfo                 `json:"web_app,omitempty"`
}

type KeyboardButtonRequestUsers struct {
	RequestID       int  `json:"request_id"`
	UserIsBot       bool `json:"user_is_bot,omitempty"`
	UserIsPremium   bool `json:"user_is_premium,omitempty"`
	MaxQuantity     int  `json:"max_quantity,omitempty"`
	RequestName     bool `json:"request_name,omitempty"`
	RequestUsername bool `json:"request_username,omitempty"`
	RequestPhoto    bool `json:"request_photo,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestID               int                      `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
	RequestTitle            bool                     `json:"request_title,omitempty"`
	RequestUsername         bool                     `json:"request_username,omitempty"`
	RequestPhoto            bool                     `json:"request_photo,omitempty"`
}

type KeyboardButtonPollType struct {
	Type enums.PollType `json:"type,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // Всегда True
	Selective      bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	IconCustomEmojiID            string                       `json:"icon_custom_emoji_id,omitempty"`
	Style                        enums.KeyboardButtonStyle    `json:"style,omitempty"`
	URL                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

type LoginUrl struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type CallbackQuery struct {
	ID              string                    `json:"id"`
	From            User                      `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"-"`
	InlineMessageID string                    `json:"inline_message_id,omitempty"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            string                    `json:"data,omitempty"`
	GameShortName   string                    `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"` // Всегда True
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       int    `json:"subscription_price,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous             bool `json:"is_anonymous"`
	CanManageChat           bool `json:"can_manage_chat"`
	CanDeleteMessages       bool `json:"can_delete_messages"`
	CanManageVideoChats     bool `json:"can_manage_video_chats"`
	CanRestrictMembers      bool `json:"can_restrict_members"`
	CanPromoteMembers       bool `json:"can_promote_members"`
	CanChangeInfo           bool `json:"can_change_info"`
	CanInviteUsers          bool `json:"can_invite_users"`
	CanPostStories          bool `json:"can_post_stories,omitempty"`
	CanEditStories          bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool `json:"can_post_messages,omitempty"`
	CanEditMessages         bool `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool `json:"can_manage_direct_messages,omitempty"`
	CanManageTags           bool `json:"can_manage_tags,omitempty"`
}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           ChatMember      `json:"-"`
	NewChatMember           ChatMember      `json:"-"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

func (c *ChatMemberUpdated) UnmarshalJSON(data []byte) error {
	type Alias ChatMemberUpdated
	aux := &struct {
		*Alias
		OldChatMember json.RawMessage `json:"old_chat_member"`
		NewChatMember json.RawMessage `json:"new_chat_member"`
	}{Alias: (*Alias)(c)}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	var err error
	if c.OldChatMember, err = unmarshalChatMember(aux.OldChatMember); err != nil {
		return err
	}
	if c.NewChatMember, err = unmarshalChatMember(aux.NewChatMember); err != nil {
		return err
	}
	return nil
}

func unmarshalChatMember(data json.RawMessage) (ChatMember, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Status enums.ChatMemberStatus `json:"status"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Status {
	case enums.MemberCreator:
		var v ChatMemberOwner
		return v, json.Unmarshal(data, &v)
	case enums.MemberAdministrator:
		var v ChatMemberAdministrator
		return v, json.Unmarshal(data, &v)
	case enums.MemberMember:
		var v ChatMemberMember
		return v, json.Unmarshal(data, &v)
	case enums.MemberRestricted:
		var v ChatMemberRestricted
		return v, json.Unmarshal(data, &v)
	case enums.MemberLeft:
		var v ChatMemberLeft
		return v, json.Unmarshal(data, &v)
	case enums.MemberKicked:
		var v ChatMemberBanned
		return v, json.Unmarshal(data, &v)
	default:
		var v ChatMemberMember
		return v, json.Unmarshal(data, &v)
	}
}

type ChatMember interface {
	isChatMember()
}

func (ChatMemberOwner) isChatMember()         {}
func (ChatMemberAdministrator) isChatMember() {}
func (ChatMemberMember) isChatMember()        {}
func (ChatMemberRestricted) isChatMember()    {}
func (ChatMemberLeft) isChatMember()          {}
func (ChatMemberBanned) isChatMember()        {}

type ChatMemberOwner struct {
	Status      enums.ChatMemberStatus `json:"status"`
	User        User                   `json:"user"`
	IsAnonymous bool                   `json:"is_anonymous"`
	CustomTitle string                 `json:"custom_title,omitempty"`
}

type ChatMemberAdministrator struct {
	Status                  enums.ChatMemberStatus `json:"status"`
	User                    User                   `json:"user"`
	CanBeEdited             bool                   `json:"can_be_edited"`
	IsAnonymous             bool                   `json:"is_anonymous"`
	CanManageChat           bool                   `json:"can_manage_chat"`
	CanDeleteMessages       bool                   `json:"can_delete_messages"`
	CanManageVideoChats     bool                   `json:"can_manage_video_chats"`
	CanRestrictMembers      bool                   `json:"can_restrict_members"`
	CanPromoteMembers       bool                   `json:"can_promote_members"`
	CanChangeInfo           bool                   `json:"can_change_info"`
	CanInviteUsers          bool                   `json:"can_invite_users"`
	CanPostStories          bool                   `json:"can_post_stories,omitempty"`
	CanEditStories          bool                   `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool                   `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool                   `json:"can_post_messages,omitempty"`
	CanEditMessages         bool                   `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool                   `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool                   `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool                   `json:"can_manage_direct_messages,omitempty"`
	CustomTitle             string                 `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
	Status    enums.ChatMemberStatus `json:"status"`
	User      User                   `json:"user"`
	Tag       string                 `json:"tag,omitempty"`
	UntilDate int64                  `json:"until_date,omitempty"`
}

type ChatMemberRestricted struct {
	Status                enums.ChatMemberStatus `json:"status"`
	User                  User                   `json:"user"`
	Tag                   string                 `json:"tag,omitempty"`
	IsMember              bool                   `json:"is_member"`
	CanSendMessages       bool                   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool                   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool                   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool                   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool                   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool                   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool                   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool                   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool                   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool                   `json:"can_add_web_page_previews,omitempty"`
	CanEditTag            bool                   `json:"can_edit_tag,omitempty"`
	CanChangeInfo         bool                   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool                   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool                   `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool                   `json:"can_manage_topics,omitempty"`
	UntilDate             int64                  `json:"until_date"`
}

type ChatMemberLeft struct {
	Status enums.ChatMemberStatus `json:"status"`
	User   User                   `json:"user"`
}

type ChatMemberBanned struct {
	Status    enums.ChatMemberStatus `json:"status"`
	User      User                   `json:"user"`
	UntilDate int64                  `json:"until_date"`
}

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanEditTag            bool `json:"can_edit_tag,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year,omitempty"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"` // Note: Sticker struct not fully defined here
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type UserRating struct {
	Level              int64 `json:"level"`
	Rating             int64 `json:"rating"`
	CurrentLevelRating int64 `json:"current_level_rating"`
	NextLevelRating    int64 `json:"next_level_rating"`
}

type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage,omitempty"`
}

type LocationAddress struct {
	CountryCode string `json:"country_code,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
}

type StoryAreaType interface {
	isStoryAreaType()
}

func (StoryAreaTypeLocation) isStoryAreaType()          {}
func (StoryAreaTypeSuggestedReaction) isStoryAreaType() {}
func (StoryAreaTypeLink) isStoryAreaType()              {}
func (StoryAreaTypeWeather) isStoryAreaType()           {}
func (StoryAreaTypeUniqueGift) isStoryAreaType()        {}

type StoryAreaTypeLocation struct {
	Type      enums.StoryAreaType `json:"type"`
	Latitude  float64             `json:"latitude"`
	Longitude float64             `json:"longitude"`
	Address   LocationAddress     `json:"address"`
}

type StoryAreaTypeSuggestedReaction struct {
	Type         enums.StoryAreaType `json:"type"`
	ReactionType ReactionType        `json:"-"`
	IsDark       bool                `json:"is_dark,omitempty"`
	IsFlipped    bool                `json:"is_flipped,omitempty"`
}

type StoryAreaTypeLink struct {
	Type enums.StoryAreaType `json:"type"`
	URL  string              `json:"url"`
}

type StoryAreaTypeWeather struct {
	Type            enums.StoryAreaType `json:"type"`
	Temperature     float64             `json:"temperature"`
	Emoji           string              `json:"emoji"`
	BackgroundColor int                 `json:"background_color"`
}

type StoryAreaTypeUniqueGift struct {
	Type enums.StoryAreaType `json:"type"`
	Name string              `json:"name"`
}

type StoryArea struct {
	Position StoryAreaPosition `json:"position"`
	Type     StoryAreaType     `json:"-"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

type ReactionType interface {
	isReactionType()
}

func (ReactionTypeEmoji) isReactionType()       {}
func (ReactionTypeCustomEmoji) isReactionType() {}
func (ReactionTypePaid) isReactionType()        {}

type ReactionTypeEmoji struct {
	Type  enums.ReactionType `json:"type"`
	Emoji string             `json:"emoji"`
}

type ReactionTypeCustomEmoji struct {
	Type          enums.ReactionType `json:"type"`
	CustomEmojiID string             `json:"custom_emoji_id"`
}

type ReactionTypePaid struct {
	Type enums.ReactionType `json:"type"`
}

type ReactionCount struct {
	Type       ReactionType `json:"-"`
	TotalCount int          `json:"total_count"`
}

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"-"`
	NewReaction []ReactionType `json:"-"`
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

type Gift struct {
	ID                     string          `json:"id"`
	Sticker                *Sticker        `json:"sticker"`
	StarCount              int64           `json:"star_count"`
	UpgradeStarCount       int64           `json:"upgrade_star_count,omitempty"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

type UniqueGiftModel struct {
	Name           string                      `json:"name"`
	Sticker        *Sticker                    `json:"sticker,omitempty"`
	RarityPerMille int                         `json:"rarity_per_mille"`
	Rarity         enums.UniqueGiftModelRarity `json:"rarity,omitempty"`
}

type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker,omitempty"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors"`
}

type UniqueGift struct {
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Number           int                `json:"number"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool               `json:"is_premium,omitempty"`
	IsBurned         bool               `json:"is_burned,omitempty"`
	IsFromBlockchain bool               `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors  `json:"colors,omitempty"`
	PublisherChat    *Chat              `json:"publisher_chat,omitempty"`
}

type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int64           `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int64           `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	UniqueGiftNumber        int64           `json:"unique_gift_number,omitempty"`
}

type UniqueGiftInfo struct {
	Gift               UniqueGift                 `json:"gift"`
	Origin             enums.UniqueGiftInfoOrigin `json:"origin"`
	LastResaleCurrency string                     `json:"last_resale_currency,omitempty"`
	LastResaleAmount   int                        `json:"last_resale_amount,omitempty"`
	OwnedGiftID        string                     `json:"owned_gift_id,omitempty"`
	TransferStarCount  int64                      `json:"transfer_star_count,omitempty"`
	NextTransferDate   int64                      `json:"next_transfer_date,omitempty"`
}

type OwnedGift interface {
	isOwnedGift()
}

func (OwnedGiftRegular) isOwnedGift() {}
func (OwnedGiftUnique) isOwnedGift()  {}

type OwnedGiftRegular struct {
	Type                    enums.OwnedGiftType `json:"type"`
	Gift                    Gift                `json:"gift"`
	OwnedGiftID             string              `json:"owned_gift_id,omitempty"`
	SenderUser              User                `json:"sender_user,omitempty"`
	SendDate                int64               `json:"send_date,omitempty"`
	Text                    string              `json:"text,omitempty"`
	Entities                []MessageEntity     `json:"entities,omitempty"`
	IsPrivate               bool                `json:"is_private,omitempty"`
	IsSaved                 bool                `json:"is_saved,omitempty"`
	CanBeUpgraded           bool                `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool                `json:"was_refunded,omitempty"`
	ConvertStarCount        int64               `json:"convert_star_count"`
	PrepaidUpgradeStarCount int64               `json:"prepaid_upgrade_star_count"`
	IsUpgradeSeparate       bool                `json:"is_upgrade_separate,omitempty"`
	UniqueGiftNumber        int64               `json:"unique_gift_number,omitempty"`
}

type OwnedGiftUnique struct {
	Type              enums.OwnedGiftType `json:"type"`
	Gift              UniqueGift          `json:"gift"`
	OwnedGiftID       string              `json:"owned_gift_id,omitempty"`
	SenderUser        User                `json:"sender_user,omitempty"`
	SendDate          int64               `json:"send_date,omitempty"`
	IsSaved           bool                `json:"is_saved,omitempty"`
	CanBeTransferred  bool                `json:"can_be_transferred,omitempty"`
	TransferStarCount int64               `json:"transfer_star_count,omitempty"`
	NextTransferDate  int64               `json:"next_transfer_date,omitempty"`
}

type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"-"`
	NextOffset string      `json:"next_offset,omitempty"`
}

type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts,omitempty"`
	LimitedGifts        bool `json:"limited_gifts,omitempty"`
	UniqueGifts         bool `json:"unique_gifts,omitempty"`
	PremiumSubscription bool `json:"premium_subscription,omitempty"`
	GiftsFromChannels   bool `json:"gifts_from_channels,omitempty"`
}

type StarAmount struct {
	Amount         int64 `json:"amount"`
	NanostarAmount int64 `json:"nanostar_amount"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope interface {
	isBotCommandScope()
}

func (BotCommandScopeDefault) isBotCommandScope()               {}
func (BotCommandScopeAllPrivateChats) isBotCommandScope()       {}
func (BotCommandScopeAllGroupChats) isBotCommandScope()         {}
func (BotCommandScopeAllChatAdministrators) isBotCommandScope() {}
func (BotCommandScopeChat) isBotCommandScope()                  {}
func (BotCommandScopeChatAdministrators) isBotCommandScope()    {}
func (BotCommandScopeChatMember) isBotCommandScope()            {}

type BotCommandScopeDefault struct {
	Type enums.BotCommandScopeType `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	Type enums.BotCommandScopeType `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	Type enums.BotCommandScopeType `json:"type"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type enums.BotCommandScopeType `json:"type"`
}

type BotCommandScopeChat struct {
	Type   enums.BotCommandScopeType `json:"type"`
	ChatID interface{}               `json:"chat_id"`
}

type BotCommandScopeChatAdministrators struct {
	Type   enums.BotCommandScopeType `json:"type"`
	ChatID interface{}               `json:"chat_id"`
}

type BotCommandScopeChatMember struct {
	Type   enums.BotCommandScopeType `json:"type"`
	ChatID interface{}               `json:"chat_id"`
	UserID int64                     `json:"user_id"`
}

type BotName struct {
	Name string `json:"name"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

type MenuButton interface {
	isMenuButton()
}

func (MenuButtonCommands) isMenuButton() {}
func (MenuButtonWebApp) isMenuButton()   {}
func (MenuButtonDefault) isMenuButton()  {}

type MenuButtonCommands struct {
	Type enums.MenuButtonType `json:"type"`
}

type MenuButtonWebApp struct {
	Type   enums.MenuButtonType `json:"type"`
	Text   string               `json:"text"`
	WebApp WebAppInfo           `json:"web_app"`
}

type MenuButtonDefault struct {
	Type enums.MenuButtonType `json:"type"`
}

type ChatBoostSource interface {
	isChatBoostSource()
}

func (ChatBoostSourcePremium) isChatBoostSource()  {}
func (ChatBoostSourceGiftCode) isChatBoostSource() {}
func (ChatBoostSourceGiveaway) isChatBoostSource() {}

type ChatBoostSourcePremium struct {
	Source enums.ChatBoostSource `json:"source"`
	User   User                  `json:"user"`
}

type ChatBoostSourceGiftCode struct {
	Source enums.ChatBoostSource `json:"source"`
	User   User                  `json:"user"`
}

type ChatBoostSourceGiveaway struct {
	Source            enums.ChatBoostSource `json:"source"`
	GiveawayMessageID int                   `json:"giveaway_message_id"`
	User              *User                 `json:"user,omitempty"`
	PrizeStarCount    int64                 `json:"prize_star_count,omitempty"`
	IsUnclaimed       bool                  `json:"is_unclaimed,omitempty"`
}

type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int64           `json:"add_date"`
	ExpirationDate int64           `json:"expiration_date"`
	Source         ChatBoostSource `json:"-"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int64           `json:"remove_date"`
	Source     ChatBoostSource `json:"-"`
}

type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner,omitempty"`
}

type ChatOwnerChanged struct {
	NewOwner User `json:"new_owner"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply,omitempty"`
	CanReadMessages            bool `json:"can_read_messages,omitempty"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages,omitempty"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                bool `json:"can_edit_name,omitempty"`
	CanEditBio                 bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           bool `json:"can_manage_stories,omitempty"`
}

type BusinessConnection struct {
	ID         string            `json:"id"`
	User       User              `json:"user"`
	UserChatID int64             `json:"user_chat_id"`
	Date       int64             `json:"date"`
	Rights     BusinessBotRights `json:"rights,omitempty"`
	IsEnabled  bool              `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}

type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

type InputMedia interface {
	isInputMedia()
}

func (InputMediaAnimation) isInputMedia() {}
func (InputMediaDocument) isInputMedia()  {}
func (InputMediaAudio) isInputMedia()     {}
func (InputMediaPhoto) isInputMedia()     {}
func (InputMediaVideo) isInputMedia()     {}

type InputMediaGroup interface {
	isInputMediaGroup()
}

func (InputMediaDocument) isInputMediaGroup() {}
func (InputMediaAudio) isInputMediaGroup()    {}
func (InputMediaPhoto) isInputMediaGroup()    {}
func (InputMediaVideo) isInputMediaGroup()    {}

type InputMediaPhoto struct {
	Type                  enums.InputMediaType `json:"type"`
	Media                 string               `json:"media"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                 `json:"has_spoiler,omitempty"`
}

type InputMediaVideo struct {
	Type                  enums.InputMediaType `json:"type"`
	Media                 string               `json:"media"`
	Thumbnail             *InputFile           `json:"thumbnail,omitempty"`
	Cover                 *InputFile           `json:"cover,omitempty"`
	StartTimestamp        int                  `json:"start_timestamp,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	Width                 int                  `json:"width,omitempty"`
	Height                int                  `json:"height,omitempty"`
	Duration              int                  `json:"duration,omitempty"`
	SupportsStreaming     bool                 `json:"supports_streaming,omitempty"`
	HasSpoiler            bool                 `json:"has_spoiler,omitempty"`
}

type InputMediaAnimation struct {
	Type                  enums.InputMediaType `json:"type"`
	Media                 string               `json:"media"`
	Thumbnail             *InputFile           `json:"thumbnail,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	Width                 int                  `json:"width,omitempty"`
	Height                int                  `json:"height,omitempty"`
	Duration              int                  `json:"duration,omitempty"`
	HasSpoiler            bool                 `json:"has_spoiler,omitempty"`
}

type InputMediaAudio struct {
	Type                  enums.InputMediaType `json:"type"`
	Media                 string               `json:"media"`
	Thumbnail             *InputFile           `json:"thumbnail,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	Duration              int                  `json:"duration,omitempty"`
	Performer             string               `json:"performer,omitempty"`
	Title                 string               `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        enums.InputMediaType `json:"type"`
	Media                       string               `json:"media"`
	Thumbnail                   *InputFile           `json:"thumbnail,omitempty"`
	Caption                     string               `json:"caption,omitempty"`
	ParseMode                   string               `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity      `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                 `json:"disable_content_type_detection,omitempty"`
}

type InputFile struct{}

type InputPaidMedia interface {
	isInputPaidMedia()
}

func (InputPaidMediaPhoto) isInputPaidMedia() {}
func (InputPaidMediaVideo) isInputPaidMedia() {}

type InputPaidMediaPhoto struct {
	Type  enums.InputPaidMediaType `json:"type"`
	Media string                   `json:"media"`
}

type InputPaidMediaVideo struct {
	Type              enums.InputPaidMediaType `json:"type"`
	Media             string                   `json:"media"`
	Thumbnail         *InputFile               `json:"thumbnail,omitempty"`
	Cover             *InputFile               `json:"cover,omitempty"`
	StartTimestamp    int                      `json:"start_timestamp,omitempty"`
	Width             int                      `json:"width,omitempty"`
	Height            int                      `json:"height,omitempty"`
	Duration          int                      `json:"duration,omitempty"`
	SupportsStreaming bool                     `json:"supports_streaming,omitempty"`
}

type InputProfilePhoto interface {
	isInputProfilePhoto()
}

func (InputProfilePhotoStatic) isInputProfilePhoto()   {}
func (InputProfilePhotoAnimated) isInputProfilePhoto() {}

type InputProfilePhotoStatic struct {
	Type  enums.InputProfilePhotoType `json:"type"`
	Photo string                      `json:"photo"`
}

type InputProfilePhotoAnimated struct {
	Type               enums.InputProfilePhotoType `json:"type"`
	Animation          string                      `json:"animation"`
	MainFrameTimestamp float64                     `json:"main_frame_timestamp,omitempty"`
}

type InputStoryContent interface {
	isInputStoryContent()
}

func (InputStoryContentPhoto) isInputStoryContent() {}
func (InputStoryContentVideo) isInputStoryContent() {}

type InputStoryContentPhoto struct {
	Type  enums.InputStoryContentType `json:"type"`
	Photo string                      `json:"photo"`
}

type InputStoryContentVideo struct {
	Type                enums.InputStoryContentType `json:"type"`
	Video               string                      `json:"video"`
	Duration            float64                     `json:"duration,omitempty"`
	CoverFrameTimestamp float64                     `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         bool                        `json:"is_animation,omitempty"`
}

// --- Unmarshal helpers ---

func unmarshalMessageOrigin(data json.RawMessage) (MessageOrigin, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.MessageOriginType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.OriginUser:
		var v MessageOriginUser
		return v, json.Unmarshal(data, &v)
	case enums.OriginHiddenUser:
		var v MessageOriginHiddenUser
		return v, json.Unmarshal(data, &v)
	case enums.OriginChat:
		var v MessageOriginChat
		return v, json.Unmarshal(data, &v)
	case enums.OriginChannel:
		var v MessageOriginChannel
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalMaybeInaccessibleMessage(data json.RawMessage) (MaybeInaccessibleMessage, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Date int64 `json:"date"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	if probe.Date == 0 {
		var v InaccessibleMessage
		return v, json.Unmarshal(data, &v)
	}
	var v Message
	return v, json.Unmarshal(data, &v)
}

func unmarshalReactionType(data json.RawMessage) (ReactionType, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.ReactionType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.ReactionEmoji:
		var v ReactionTypeEmoji
		return v, json.Unmarshal(data, &v)
	case enums.ReactionCustomEmoji:
		var v ReactionTypeCustomEmoji
		return v, json.Unmarshal(data, &v)
	case enums.ReactionPaid:
		var v ReactionTypePaid
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalReactionTypeSlice(data []json.RawMessage) ([]ReactionType, error) {
	if len(data) == 0 {
		return nil, nil
	}
	result := make([]ReactionType, 0, len(data))
	for _, raw := range data {
		rt, err := unmarshalReactionType(raw)
		if err != nil {
			return nil, err
		}
		if rt != nil {
			result = append(result, rt)
		}
	}
	return result, nil
}

func unmarshalChatBoostSource(data json.RawMessage) (ChatBoostSource, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Source enums.ChatBoostSource `json:"source"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Source {
	case enums.BoostPremium:
		var v ChatBoostSourcePremium
		return v, json.Unmarshal(data, &v)
	case enums.BoostGiftCode:
		var v ChatBoostSourceGiftCode
		return v, json.Unmarshal(data, &v)
	case enums.BoostGiveaway:
		var v ChatBoostSourceGiveaway
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalPaidMedia(data json.RawMessage) (PaidMedia, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.PaidMediaType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.PaidMediaPreview:
		var v PaidMediaPreview
		return v, json.Unmarshal(data, &v)
	case enums.PaidMediaPhoto:
		var v PaidMediaPhoto
		return v, json.Unmarshal(data, &v)
	case enums.PaidMediaVideo:
		var v PaidMediaVideo
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalPaidMediaSlice(data []json.RawMessage) ([]PaidMedia, error) {
	if len(data) == 0 {
		return nil, nil
	}
	result := make([]PaidMedia, 0, len(data))
	for _, raw := range data {
		pm, err := unmarshalPaidMedia(raw)
		if err != nil {
			return nil, err
		}
		if pm != nil {
			result = append(result, pm)
		}
	}
	return result, nil
}

func unmarshalBackgroundFill(data json.RawMessage) (BackgroundFill, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.BackgroundFillType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.FillSolid:
		var v BackgroundFillSolid
		return v, json.Unmarshal(data, &v)
	case enums.FillGradient:
		var v BackgroundFillGradient
		return v, json.Unmarshal(data, &v)
	case enums.FillFreeformGradient:
		var v BackgroundFillFreeformGradient
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalBackgroundType(data json.RawMessage) (BackgroundType, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.BackgroundType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.BackgroundFill:
		var v BackgroundTypeFill
		return v, json.Unmarshal(data, &v)
	case enums.BackgroundWallpaper:
		var v BackgroundTypeWallpaper
		return v, json.Unmarshal(data, &v)
	case enums.BackgroundPattern:
		var v BackgroundTypePattern
		return v, json.Unmarshal(data, &v)
	case enums.BackgroundChatTheme:
		var v BackgroundTypeChatTheme
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalStoryAreaType(data json.RawMessage) (StoryAreaType, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.StoryAreaType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.StoryAreaLocation:
		var v StoryAreaTypeLocation
		return v, json.Unmarshal(data, &v)
	case enums.StoryAreaSuggestedReaction:
		var v StoryAreaTypeSuggestedReaction
		return v, json.Unmarshal(data, &v)
	case enums.StoryAreaLink:
		var v StoryAreaTypeLink
		return v, json.Unmarshal(data, &v)
	case enums.StoryAreaWeather:
		var v StoryAreaTypeWeather
		return v, json.Unmarshal(data, &v)
	case enums.StoryAreaUniqueGift:
		var v StoryAreaTypeUniqueGift
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalOwnedGift(data json.RawMessage) (OwnedGift, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.OwnedGiftType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.OwnedGiftRegular:
		var v OwnedGiftRegular
		return v, json.Unmarshal(data, &v)
	case enums.OwnedGiftUnique:
		var v OwnedGiftUnique
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

// UnmarshalMenuButton decodes a MenuButton from JSON by discriminating on the "type" field.
func UnmarshalMenuButton(data json.RawMessage) (MenuButton, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.MenuButtonType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.MenuCommands:
		var v MenuButtonCommands
		return v, json.Unmarshal(data, &v)
	case enums.MenuWebApp:
		var v MenuButtonWebApp
		return v, json.Unmarshal(data, &v)
	case enums.MenuDefault:
		var v MenuButtonDefault
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

// --- UnmarshalJSON methods ---

func (e *ExternalReplyInfo) UnmarshalJSON(data []byte) error {
	type Alias ExternalReplyInfo
	aux := &struct {
		*Alias
		Origin json.RawMessage `json:"origin"`
	}{Alias: (*Alias)(e)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.Origin) > 0 {
		origin, err := unmarshalMessageOrigin(aux.Origin)
		if err != nil {
			return err
		}
		e.Origin = &origin
	}
	return nil
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		*Alias
		ForwardOrigin json.RawMessage `json:"forward_origin"`
		PinnedMessage json.RawMessage `json:"pinned_message"`
	}{Alias: (*Alias)(m)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.ForwardOrigin) > 0 {
		origin, err := unmarshalMessageOrigin(aux.ForwardOrigin)
		if err != nil {
			return err
		}
		m.ForwardOrigin = &origin
	}
	if len(aux.PinnedMessage) > 0 {
		msg, err := unmarshalMaybeInaccessibleMessage(aux.PinnedMessage)
		if err != nil {
			return err
		}
		m.PinnedMessage = &msg
	}
	return nil
}

func (c *CallbackQuery) UnmarshalJSON(data []byte) error {
	type Alias CallbackQuery
	aux := &struct {
		*Alias
		Message json.RawMessage `json:"message"`
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.Message) > 0 {
		msg, err := unmarshalMaybeInaccessibleMessage(aux.Message)
		if err != nil {
			return err
		}
		c.Message = &msg
	}
	return nil
}

func (r *ReactionCount) UnmarshalJSON(data []byte) error {
	type Alias ReactionCount
	aux := &struct {
		*Alias
		Type json.RawMessage `json:"type"`
	}{Alias: (*Alias)(r)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if r.Type, err = unmarshalReactionType(aux.Type); err != nil {
		return err
	}
	return nil
}

func (m *MessageReactionUpdated) UnmarshalJSON(data []byte) error {
	type Alias MessageReactionUpdated
	aux := &struct {
		*Alias
		OldReaction []json.RawMessage `json:"old_reaction"`
		NewReaction []json.RawMessage `json:"new_reaction"`
	}{Alias: (*Alias)(m)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if m.OldReaction, err = unmarshalReactionTypeSlice(aux.OldReaction); err != nil {
		return err
	}
	if m.NewReaction, err = unmarshalReactionTypeSlice(aux.NewReaction); err != nil {
		return err
	}
	return nil
}

func (s *StoryAreaTypeSuggestedReaction) UnmarshalJSON(data []byte) error {
	type Alias StoryAreaTypeSuggestedReaction
	aux := &struct {
		*Alias
		ReactionType json.RawMessage `json:"reaction_type"`
	}{Alias: (*Alias)(s)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if s.ReactionType, err = unmarshalReactionType(aux.ReactionType); err != nil {
		return err
	}
	return nil
}

func (p *PaidMediaInfo) UnmarshalJSON(data []byte) error {
	type Alias PaidMediaInfo
	aux := &struct {
		*Alias
		PaidMedia []json.RawMessage `json:"paid_media"`
	}{Alias: (*Alias)(p)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if p.PaidMedia, err = unmarshalPaidMediaSlice(aux.PaidMedia); err != nil {
		return err
	}
	return nil
}

func (b *BackgroundTypeFill) UnmarshalJSON(data []byte) error {
	type Alias BackgroundTypeFill
	aux := &struct {
		*Alias
		Fill json.RawMessage `json:"fill"`
	}{Alias: (*Alias)(b)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if b.Fill, err = unmarshalBackgroundFill(aux.Fill); err != nil {
		return err
	}
	return nil
}

func (b *BackgroundTypePattern) UnmarshalJSON(data []byte) error {
	type Alias BackgroundTypePattern
	aux := &struct {
		*Alias
		Fill json.RawMessage `json:"fill"`
	}{Alias: (*Alias)(b)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if b.Fill, err = unmarshalBackgroundFill(aux.Fill); err != nil {
		return err
	}
	return nil
}

func (c *ChatBackground) UnmarshalJSON(data []byte) error {
	type Alias ChatBackground
	aux := &struct {
		*Alias
		Type json.RawMessage `json:"type"`
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if c.Type, err = unmarshalBackgroundType(aux.Type); err != nil {
		return err
	}
	return nil
}

func (s *StoryArea) UnmarshalJSON(data []byte) error {
	type Alias StoryArea
	aux := &struct {
		*Alias
		Type json.RawMessage `json:"type"`
	}{Alias: (*Alias)(s)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if s.Type, err = unmarshalStoryAreaType(aux.Type); err != nil {
		return err
	}
	return nil
}

func (o *OwnedGifts) UnmarshalJSON(data []byte) error {
	type Alias OwnedGifts
	aux := &struct {
		*Alias
		Gifts []json.RawMessage `json:"gifts"`
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.Gifts) > 0 {
		o.Gifts = make([]OwnedGift, 0, len(aux.Gifts))
		for _, raw := range aux.Gifts {
			gift, err := unmarshalOwnedGift(raw)
			if err != nil {
				return err
			}
			if gift != nil {
				o.Gifts = append(o.Gifts, gift)
			}
		}
	}
	return nil
}

func (c *ChatBoost) UnmarshalJSON(data []byte) error {
	type Alias ChatBoost
	aux := &struct {
		*Alias
		Source json.RawMessage `json:"source"`
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if c.Source, err = unmarshalChatBoostSource(aux.Source); err != nil {
		return err
	}
	return nil
}

func (c *ChatBoostRemoved) UnmarshalJSON(data []byte) error {
	type Alias ChatBoostRemoved
	aux := &struct {
		*Alias
		Source json.RawMessage `json:"source"`
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if c.Source, err = unmarshalChatBoostSource(aux.Source); err != nil {
		return err
	}
	return nil
}

package enums

type (
	AllowedUpdates        string
	ChatType              string
	MessageEntityType     string
	MessageOriginType     string
	PaidMediaType         string
	DiceValue             string
	PollType              string
	VenueFoursquareType   string
	BackgroundFillType    string
	BackgroundType        string
	KeyboardButtonStyle   string
	ChatMemberStatus      string
	StoryAreaType         string
	ReactionType          string
	UniqueGiftModelRarity string
	UniqueGiftInfoOrigin  string
	OwnedGiftType         string
	ChatBoostSource       string
	VideoQualityCodec          string
	StickerType                string
	StickerFormat              string
	MaskPositionPoint          string
	InputMediaType             string
	InputPaidMediaType         string
	InputProfilePhotoType      string
	InputStoryContentType      string
	BotCommandScopeType        string
	MenuButtonType             string
	InlineQueryResultType      string
	PassportElementType        string
	RevenueWithdrawalStateType string
	TransactionPartnerType     string
)

// AllowedUpdates
const (
	UpdateMessage              AllowedUpdates = "message"
	UpdateEditedMessage        AllowedUpdates = "edited_message"
	UpdateChannelPost          AllowedUpdates = "channel_post"
	UpdateEditedChannelPost    AllowedUpdates = "edited_channel_post"
	UpdateCallbackQuery        AllowedUpdates = "callback_query"
	UpdateInlineQuery          AllowedUpdates = "inline_query"
	UpdateChosenInlineResult   AllowedUpdates = "chosen_inline_result"
	UpdateShippingQuery        AllowedUpdates = "shipping_query"
	UpdatePreCheckoutQuery     AllowedUpdates = "pre_checkout_query"
	UpdatePoll                 AllowedUpdates = "poll"
	UpdatePollAnswer           AllowedUpdates = "poll_answer"
	UpdateMyChatMember         AllowedUpdates = "my_chat_member"
	UpdateChatMember           AllowedUpdates = "chat_member"
	UpdateChatJoinRequest      AllowedUpdates = "chat_join_request"
	UpdateChatBoost            AllowedUpdates = "chat_boost"
	UpdateRemovedChatBoost     AllowedUpdates = "removed_chat_boost"
	UpdateMessageReaction      AllowedUpdates = "message_reaction"
	UpdateMessageReactionCount AllowedUpdates = "message_reaction_count"
	UpdateBusinessConnection   AllowedUpdates = "business_connection"
	UpdateBusinessMessage      AllowedUpdates = "business_message"
	UpdatePurchasedPaidMedia   AllowedUpdates = "purchased_paid_media"
)

// ChatType
const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// MessageEntityType
const (
	EntityMention              MessageEntityType = "mention"
	EntityHashtag              MessageEntityType = "hashtag"
	EntityCashtag              MessageEntityType = "cashtag"
	EntityBotCommand           MessageEntityType = "bot_command"
	EntityURL                  MessageEntityType = "url"
	EntityEmail                MessageEntityType = "email"
	EntityPhoneNumber          MessageEntityType = "phone_number"
	EntityBold                 MessageEntityType = "bold"
	EntityItalic               MessageEntityType = "italic"
	EntityUnderline            MessageEntityType = "underline"
	EntityStrikethrough        MessageEntityType = "strikethrough"
	EntitySpoiler              MessageEntityType = "spoiler"
	EntityBlockquote           MessageEntityType = "blockquote"
	EntityExpandableBlockquote MessageEntityType = "expandable_blockquote"
	EntityPre                  MessageEntityType = "pre"
	EntityTextLink             MessageEntityType = "text_link"
	EntityTextMention          MessageEntityType = "text_mention"
	EntityCode                 MessageEntityType = "code"
	EntityCustomEmoji          MessageEntityType = "custom_emoji"
	EntityDateTime             MessageEntityType = "date_time"
)

// MessageOriginType
const (
	OriginUser       MessageOriginType = "user"
	OriginHiddenUser MessageOriginType = "hidden_user"
	OriginChat       MessageOriginType = "chat"
	OriginChannel    MessageOriginType = "channel"
)

// PaidMediaType
const (
	PaidMediaPreview PaidMediaType = "preview"
	PaidMediaPhoto   PaidMediaType = "photo"
	PaidMediaVideo   PaidMediaType = "video"
)

// DiceValue
const (
	DiceGameDie     DiceValue = "🎲"
	DiceDirectHit   DiceValue = "🎯"
	DiceBowling     DiceValue = "🎳"
	DiceBasketball  DiceValue = "🏀"
	DiceSoccerBall  DiceValue = "⚽"
	DiceSlotMachine DiceValue = "🎰"
)

// PollType
const (
	PollRegular PollType = "regular"
	PollQuiz    PollType = "quiz"
)

// VenueFoursquareType
const (
	FoursquareArtsDefault  VenueFoursquareType = "arts_entertainment/default"
	FoursquareArtsAquarium VenueFoursquareType = "arts_entertainment/aquarium"
	FoursquareFoodIcecream VenueFoursquareType = "food/icecream"
)

// BackgroundFillType
const (
	FillSolid            BackgroundFillType = "solid"
	FillGradient         BackgroundFillType = "gradient"
	FillFreeformGradient BackgroundFillType = "freeform_gradient"
)

// BackgroundType
const (
	BackgroundFill      BackgroundType = "fill"
	BackgroundWallpaper BackgroundType = "wallpaper"
	BackgroundPattern   BackgroundType = "pattern"
	BackgroundChatTheme BackgroundType = "chat_theme"
)

// KeyboardButtonStyle
const (
	ButtonDanger  KeyboardButtonStyle = "danger"
	ButtonSuccess KeyboardButtonStyle = "success"
	ButtonPrimary KeyboardButtonStyle = "primary"
)

// ChatMemberStatus
const (
	MemberCreator       ChatMemberStatus = "creator"
	MemberAdministrator ChatMemberStatus = "administrator"
	MemberMember        ChatMemberStatus = "member"
	MemberRestricted    ChatMemberStatus = "restricted"
	MemberLeft          ChatMemberStatus = "left"
	MemberKicked        ChatMemberStatus = "kicked"
)

// StoryAreaType
const (
	StoryAreaLocation          StoryAreaType = "location"
	StoryAreaSuggestedReaction StoryAreaType = "suggested_reaction"
	StoryAreaLink              StoryAreaType = "link"
	StoryAreaWeather           StoryAreaType = "weather"
	StoryAreaUniqueGift        StoryAreaType = "unique_gift"
)

// ReactionType
const (
	ReactionEmoji       ReactionType = "emoji"
	ReactionCustomEmoji ReactionType = "custom_emoji"
	ReactionPaid        ReactionType = "paid"
)

// UniqueGiftModelRarity
const (
	RarityUncommon  UniqueGiftModelRarity = "uncommon"
	RarityRare      UniqueGiftModelRarity = "rare"
	RarityEpic      UniqueGiftModelRarity = "epic"
	RarityLegendary UniqueGiftModelRarity = "legendary"
)

// UniqueGiftInfoOrigin
const (
	GiftOriginUpgrade       UniqueGiftInfoOrigin = "upgrade"
	GiftOriginTransfer      UniqueGiftInfoOrigin = "transfer"
	GiftOriginResale        UniqueGiftInfoOrigin = "resale"
	GiftOriginGiftedUpgrade UniqueGiftInfoOrigin = "gifted_upgrade"
	GiftOriginOffer         UniqueGiftInfoOrigin = "offer"
)

// OwnedGiftType
const (
	OwnedGiftRegular OwnedGiftType = "regular"
	OwnedGiftUnique  OwnedGiftType = "unique"
)

// ChatBoostSource
const (
	BoostPremium  ChatBoostSource = "premium"
	BoostGiftCode ChatBoostSource = "gift_code"
	BoostGiveaway ChatBoostSource = "giveaway"
)

// VideoQualityCodec
const (
	VideoCodecH264 VideoQualityCodec = "h264"
	VideoCodecH265 VideoQualityCodec = "h265"
	VideoCodecAV01 VideoQualityCodec = "av01"
)

// StickerType
const (
	StickerRegular     StickerType = "regular"
	StickerMask        StickerType = "mask"
	StickerCustomEmoji StickerType = "custom_emoji"
)

// StickerFormat
const (
	StickerFormatStatic   StickerFormat = "static"
	StickerFormatAnimated StickerFormat = "animated"
	StickerFormatVideo    StickerFormat = "video"
)

// MaskPositionPoint
const (
	MaskForehead MaskPositionPoint = "forehead"
	MaskEyes     MaskPositionPoint = "eyes"
	MaskMouth    MaskPositionPoint = "mouth"
	MaskChin     MaskPositionPoint = "chin"
)

// InputMediaType
const (
	InputMediaPhoto     InputMediaType = "photo"
	InputMediaVideo     InputMediaType = "video"
	InputMediaAnimation InputMediaType = "animation"
	InputMediaAudio     InputMediaType = "audio"
	InputMediaDocument  InputMediaType = "document"
)

// InputPaidMediaType
const (
	InputPaidMediaPhoto InputPaidMediaType = "photo"
	InputPaidMediaVideo InputPaidMediaType = "video"
)

// InputProfilePhotoType
const (
	InputProfileStatic   InputProfilePhotoType = "static"
	InputProfileAnimated InputProfilePhotoType = "animated"
)

// InputStoryContentType
const (
	InputStoryPhoto InputStoryContentType = "photo"
	InputStoryVideo InputStoryContentType = "video"
)

// BotCommandScopeType
const (
	ScopeDefault               BotCommandScopeType = "default"
	ScopeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	ScopeAllGroupChats         BotCommandScopeType = "all_group_chats"
	ScopeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	ScopeChat                  BotCommandScopeType = "chat"
	ScopeChatAdministrators    BotCommandScopeType = "chat_administrators"
	ScopeChatMember            BotCommandScopeType = "chat_member"
)

// MenuButtonType
const (
	MenuCommands MenuButtonType = "commands"
	MenuWebApp   MenuButtonType = "web_app"
	MenuDefault  MenuButtonType = "default"
)

// InlineQueryResultType
const (
	InlineResultArticle  InlineQueryResultType = "article"
	InlineResultPhoto    InlineQueryResultType = "photo"
	InlineResultGif      InlineQueryResultType = "gif"
	InlineResultMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineResultVideo    InlineQueryResultType = "video"
	InlineResultAudio    InlineQueryResultType = "audio"
	InlineResultVoice    InlineQueryResultType = "voice"
	InlineResultDocument InlineQueryResultType = "document"
	InlineResultLocation InlineQueryResultType = "location"
	InlineResultVenue    InlineQueryResultType = "venue"
	InlineResultContact  InlineQueryResultType = "contact"
	InlineResultGame     InlineQueryResultType = "game"
	InlineResultSticker  InlineQueryResultType = "sticker"
)

// PassportElementType
const (
	PassportPersonalDetails       PassportElementType = "personal_details"
	PassportPassport              PassportElementType = "passport"
	PassportDriverLicense         PassportElementType = "driver_license"
	PassportIdentityCard          PassportElementType = "identity_card"
	PassportInternalPassport      PassportElementType = "internal_passport"
	PassportAddress               PassportElementType = "address"
	PassportUtilityBill           PassportElementType = "utility_bill"
	PassportBankStatement         PassportElementType = "bank_statement"
	PassportRentalAgreement       PassportElementType = "rental_agreement"
	PassportPassportRegistration  PassportElementType = "passport_registration"
	PassportTemporaryRegistration PassportElementType = "temporary_registration"
	PassportPhoneNumber           PassportElementType = "phone_number"
	PassportEmail                 PassportElementType = "email"
)

// RevenueWithdrawalStateType
const (
	WithdrawalPending   RevenueWithdrawalStateType = "pending"
	WithdrawalSucceeded RevenueWithdrawalStateType = "succeeded"
	WithdrawalFailed    RevenueWithdrawalStateType = "failed"
)

// TransactionPartnerType
const (
	TransactionPartnerUser             TransactionPartnerType = "user"
	TransactionPartnerChat             TransactionPartnerType = "chat"
	TransactionPartnerAffiliateProgram TransactionPartnerType = "affiliate_program"
	TransactionPartnerFragment         TransactionPartnerType = "fragment"
	TransactionPartnerTelegramAds      TransactionPartnerType = "telegram_ads"
	TransactionPartnerTelegramApi      TransactionPartnerType = "telegram_api"
	TransactionPartnerOther            TransactionPartnerType = "other"
)

package params

import "github.com/gram-kit/gramkit-go/models"

type SendInvoice struct {
	ChatID                    any                             `json:"chat_id"`
	MessageID                 int                             `json:"message_id"`
	MessageThreadID           int                             `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID     int                             `json:"direct_messages_topic_id,omitempty"`
	Title                     string                          `json:"title"`
	Description               string                          `json:"description"`
	Payload                   string                          `json:"payload"`
	ProviderToken             string                          `json:"provider_token,omitempty"`
	Currency                  string                          `json:"currency"`
	Prices                    []models.LabeledPrice           `json:"prices"`
	MaxTipAmount              int                             `json:"max_tip_amount,omitempty"`
	SuggestedTipAmount        int                             `json:"suggested_tip_amount,omitempty"`
	StartParameter            string                          `json:"start_parameter,omitempty"`
	ProviderData              string                          `json:"provider_data,omitempty"`
	PhotoURL                  string                          `json:"photo_url,omitempty"`
	PhotoSize                 int                             `json:"photo_size,omitempty"`
	PhotoWidth                int                             `json:"photo_width,omitempty"`
	PhotoHeight               int                             `json:"photo_height,omitempty"`
	NeedName                  bool                            `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                            `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                            `json:"need_email,omitempty"`
	NeedShippingAddress       bool                            `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                            `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                            `json:"send_email_to_provider,omitempty"`
	IsFixable                 bool                            `json:"is_fixable,omitempty"`
	DisableNotification       bool                            `json:"disable_notification,omitempty"`
	ProtectContent            bool                            `json:"protect_content,omitempty"`
	AllowPaidBroadcast        bool                            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID           string                          `json:"message_effect_id,omitempty"`
	SuggestedPostParameters   *models.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters           *models.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup               *models.ReplyMarkup             `json:"reply_markup,omitempty"`
}

type CreateInvoiceLink struct {
	BusinessConnectionID      string                `json:"business_connection_id,omitempty"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token,omitempty"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	SubscriptionPeriod        int64                 `json:"subscription_period,omitempty"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmount        any                   `json:"suggested_tip_amount,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFixable                 bool                  `json:"is_fixable,omitempty"`
}

type AnswerShippingQuery struct {
	ShippingQueryID string                  `json:"shipping_query_id"`
	Ok              bool                    `json:"ok"`
	ShippingOptions []models.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                  `json:"error_message,omitempty"`
}

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

type GetStarTransactions struct {
	Offset int64 `json:"offset"`
	Limit  int   `json:"limit"`
}

type RefundStarPayment struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

type EditUserStarSubscription struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

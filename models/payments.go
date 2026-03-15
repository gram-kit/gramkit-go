package models

import (
	"encoding/json"

	"github.com/gram-kit/gramkit-go/enums"
)

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string          `json:"name,omitempty"`
	PhoneNumber     string          `json:"phone_number,omitempty"`
	Email           string          `json:"email,omitempty"`
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int64      `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date"`
	IsRecurring                bool       `json:"is_recurring,omitempty"`
	IsFirstRecurring           bool       `json:"is_first_recurring,omitempty"`
	ShippingOptionID           string     `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id,omitempty"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id,omitempty"`
}

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int64  `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id,omitempty"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

type RevenueWithdrawalState interface {
	isRevenueWithdrawalState()
}

func (RevenueWithdrawalStatePending) isRevenueWithdrawalState()   {}
func (RevenueWithdrawalStateSucceeded) isRevenueWithdrawalState() {}
func (RevenueWithdrawalStateFailed) isRevenueWithdrawalState()    {}

type RevenueWithdrawalStatePending struct {
	Type enums.RevenueWithdrawalStateType `json:"type"`
}

type RevenueWithdrawalStateSucceeded struct {
	Type enums.RevenueWithdrawalStateType `json:"type"`
	Date int64                            `json:"date"`
	URL  string                           `json:"url"`
}

type RevenueWithdrawalStateFailed struct {
	Type enums.RevenueWithdrawalStateType `json:"type"`
}

type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMille int   `json:"commission_per_mille,omitempty"`
	Amount             int64 `json:"amount,omitempty"`
	NanostarAmount     int64 `json:"nanostar_amount,omitempty"`
}

type TransactionPartner interface {
	isTransactionPartner()
}

func (TransactionPartnerUser) isTransactionPartner()             {}
func (TransactionPartnerChat) isTransactionPartner()             {}
func (TransactionPartnerAffiliateProgram) isTransactionPartner() {}
func (TransactionPartnerFragment) isTransactionPartner()         {}
func (TransactionPartnerTelegramAds) isTransactionPartner()      {}
func (TransactionPartnerTelegramApi) isTransactionPartner()      {}
func (TransactionPartnerOther) isTransactionPartner()            {}

type TransactionPartnerUser struct {
	Type                        enums.TransactionPartnerType `json:"type"`
	TransactionType             string                       `json:"transaction_type"`
	User                        User                         `json:"user"`
	Affiliate                   *AffiliateInfo               `json:"affiliate_info,omitempty"`
	InvoicePayload              string                       `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          int64                        `json:"subscription_period,omitempty"`
	PaidMedia                   []PaidMedia                  `json:"-"`
	PaidMediaPayload            string                       `json:"paid_media_payload,omitempty"`
	Gift                        *Gift                        `json:"gift,omitempty"`
	PremiumSubscriptionDuration int64                        `json:"premium_subscription_duration,omitempty"`
}

type TransactionPartnerChat struct {
	Type enums.TransactionPartnerType `json:"type"`
	Chat Chat                         `json:"chat"`
	Gift *Gift                        `json:"gift,omitempty"`
}

type TransactionPartnerAffiliateProgram struct {
	Type              enums.TransactionPartnerType `json:"type"`
	SponsorUser       *User                        `json:"sponsor_user,omitempty"`
	CommissionPerMile int64                        `json:"commission_per_mile"`
}

type TransactionPartnerFragment struct {
	Type            enums.TransactionPartnerType `json:"type"`
	WithdrawalState *RevenueWithdrawalState      `json:"-"`
}

type TransactionPartnerTelegramAds struct {
	Type enums.TransactionPartnerType `json:"type"`
}

type TransactionPartnerTelegramApi struct {
	Type         enums.TransactionPartnerType `json:"type"`
	RequestCount int                          `json:"request_count"`
}

type TransactionPartnerOther struct {
	Type enums.TransactionPartnerType `json:"type"`
}

type StarTransaction struct {
	ID             string             `json:"id"`
	Amount         int64              `json:"amount"`
	NanostarAmount int64              `json:"nanostar_amount,omitempty"`
	Date           int64              `json:"date,omitempty"`
	Source         TransactionPartner `json:"-"`
	Receiver       TransactionPartner `json:"-"`
}

type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// --- Unmarshal helpers ---

func unmarshalRevenueWithdrawalState(data json.RawMessage) (RevenueWithdrawalState, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.RevenueWithdrawalStateType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.WithdrawalPending:
		var v RevenueWithdrawalStatePending
		return v, json.Unmarshal(data, &v)
	case enums.WithdrawalSucceeded:
		var v RevenueWithdrawalStateSucceeded
		return v, json.Unmarshal(data, &v)
	case enums.WithdrawalFailed:
		var v RevenueWithdrawalStateFailed
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

func unmarshalTransactionPartner(data json.RawMessage) (TransactionPartner, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var probe struct {
		Type enums.TransactionPartnerType `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case enums.TransactionPartnerUser:
		var v TransactionPartnerUser
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerChat:
		var v TransactionPartnerChat
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerAffiliateProgram:
		var v TransactionPartnerAffiliateProgram
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerFragment:
		var v TransactionPartnerFragment
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerTelegramAds:
		var v TransactionPartnerTelegramAds
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerTelegramApi:
		var v TransactionPartnerTelegramApi
		return v, json.Unmarshal(data, &v)
	case enums.TransactionPartnerOther:
		var v TransactionPartnerOther
		return v, json.Unmarshal(data, &v)
	default:
		return nil, nil
	}
}

// --- UnmarshalJSON methods ---

func (t *TransactionPartnerFragment) UnmarshalJSON(data []byte) error {
	type Alias TransactionPartnerFragment
	aux := &struct {
		*Alias
		WithdrawalState json.RawMessage `json:"withdrawal_state"`
	}{Alias: (*Alias)(t)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.WithdrawalState) > 0 {
		ws, err := unmarshalRevenueWithdrawalState(aux.WithdrawalState)
		if err != nil {
			return err
		}
		t.WithdrawalState = &ws
	}
	return nil
}

func (t *TransactionPartnerUser) UnmarshalJSON(data []byte) error {
	type Alias TransactionPartnerUser
	aux := &struct {
		*Alias
		PaidMedia []json.RawMessage `json:"paid_media"`
	}{Alias: (*Alias)(t)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if t.PaidMedia, err = unmarshalPaidMediaSlice(aux.PaidMedia); err != nil {
		return err
	}
	return nil
}

func (s *StarTransaction) UnmarshalJSON(data []byte) error {
	type Alias StarTransaction
	aux := &struct {
		*Alias
		Source   json.RawMessage `json:"source"`
		Receiver json.RawMessage `json:"receiver"`
	}{Alias: (*Alias)(s)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	if s.Source, err = unmarshalTransactionPartner(aux.Source); err != nil {
		return err
	}
	if s.Receiver, err = unmarshalTransactionPartner(aux.Receiver); err != nil {
		return err
	}
	return nil
}

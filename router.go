package gramkit

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/gram-kit/gramkit-go/enums"
	"github.com/gram-kit/gramkit-go/models"
)

// HandlerType defines which part of an update to match.
type HandlerType int

const (
	OnMessage                 HandlerType = iota // matches update.Message
	OnEditedMessage                              // matches update.EditedMessage
	OnChannelPost                                // matches update.ChannelPost
	OnEditedChannelPost                          // matches update.EditedChannelPost
	OnCallbackQuery                              // matches update.CallbackQuery
	OnInlineQuery                                // matches update.InlineQuery
	OnChosenInlineResult                         // matches update.ChosenInlineResult
	OnShippingQuery                              // matches update.ShippingQuery
	OnPreCheckoutQuery                           // matches update.PreCheckoutQuery
	OnPoll                                       // matches update.Poll
	OnPollAnswer                                 // matches update.PollAnswer
	OnMyChatMember                               // matches update.MyChatMember
	OnChatMember                                 // matches update.ChatMember
	OnChatJoinRequest                            // matches update.ChatJoinRequest
	OnChatBoost                                  // matches update.ChatBoost
	OnRemovedChatBoost                           // matches update.RemovedChatBoost
	OnMessageReaction                            // matches update.MessageReaction
	OnMessageReactionCount                       // matches update.MessageReactionCount
	OnBusinessConnection                         // matches update.BusinessConnection
	OnBusinessMessage                            // matches update.BusinessMessage
	OnEditedBusinessMessage                      // matches update.EditedBusinessMessage
	OnDeletedBusinessMessages                    // matches update.DeletedBusinessMessages
	OnPurchasedPaidMedia                         // matches update.PurchasedPaidMedia
	OnCommand                                    // matches bot commands (e.g. /start, /help)
)

// MatchType defines how to match the pattern.
type MatchType int

const (
	MatchExact    MatchType = iota // exact match
	MatchPrefix                    // prefix match (e.g. "/start" matches "/start@bot" and "/start hello")
	MatchContains                  // contains match
)

// handler is a registered handler with its matching criteria.
type handler struct {
	handlerType HandlerType
	pattern     string
	matchType   MatchType
	handler     HandlerFunc
}

// RegisterHandler registers a handler for a specific update type with exact match.
func (b *Bot) RegisterHandler(handlerType HandlerType, pattern string, h HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers = append(b.handlers, handler{
		handlerType: handlerType,
		pattern:     pattern,
		matchType:   MatchExact,
		handler:     h,
	})
}

// RegisterHandlerMatch registers a handler with a specific match type.
func (b *Bot) RegisterHandlerMatch(handlerType HandlerType, matchType MatchType, pattern string, h HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers = append(b.handlers, handler{
		handlerType: handlerType,
		pattern:     pattern,
		matchType:   matchType,
		handler:     h,
	})
}

// matchHandler finds the first matching handler for an update.
func (b *Bot) matchHandler(update *models.Update) HandlerFunc {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, h := range b.handlers {
		text := extractText(h.handlerType, update)
		if text == "" {
			continue
		}
		if matchPattern(h.matchType, text, h.pattern) {
			return h.handler
		}
	}

	return b.defaultHandler
}

// extractText returns the relevant text from an update based on handler type.
func extractText(ht HandlerType, u *models.Update) string {
	switch ht {
	case OnMessage:
		if u.Message != nil {
			return u.Message.Text
		}
	case OnEditedMessage:
		if u.EditedMessage != nil {
			return u.EditedMessage.Text
		}
	case OnChannelPost:
		if u.ChannelPost != nil {
			return u.ChannelPost.Text
		}
	case OnEditedChannelPost:
		if u.EditedChannelPost != nil {
			return u.EditedChannelPost.Text
		}
	case OnCallbackQuery:
		if u.CallbackQuery != nil {
			return u.CallbackQuery.Data
		}
	case OnInlineQuery:
		if u.InlineQuery != nil {
			return u.InlineQuery.Query
		}
	case OnBusinessMessage:
		if u.BusinessMessage != nil {
			return u.BusinessMessage.Text
		}
	case OnEditedBusinessMessage:
		if u.EditedBusinessMessage != nil {
			return u.EditedBusinessMessage.Text
		}
	// Types without text — return a marker so the handler fires if the update type matches.
	case OnChosenInlineResult:
		if u.ChosenInlineResult != nil {
			return u.ChosenInlineResult.Query
		}
	case OnShippingQuery:
		if u.ShippingQuery != nil {
			return "*"
		}
	case OnPreCheckoutQuery:
		if u.PreCheckoutQuery != nil {
			return "*"
		}
	case OnPoll:
		if u.Poll != nil {
			return "*"
		}
	case OnPollAnswer:
		if u.PollAnswer != nil {
			return "*"
		}
	case OnMyChatMember:
		if u.MyChatMember != nil {
			return "*"
		}
	case OnChatMember:
		if u.ChatMember != nil {
			return "*"
		}
	case OnChatJoinRequest:
		if u.ChatJoinRequest != nil {
			return "*"
		}
	case OnChatBoost:
		if u.ChatBoost != nil {
			return "*"
		}
	case OnRemovedChatBoost:
		if u.RemovedChatBoost != nil {
			return "*"
		}
	case OnMessageReaction:
		if u.MessageReaction != nil {
			return "*"
		}
	case OnMessageReactionCount:
		if u.MessageReactionCount != nil {
			return "*"
		}
	case OnBusinessConnection:
		if u.BusinessConnection != nil {
			return "*"
		}
	case OnDeletedBusinessMessages:
		if u.DeletedBusinessMessages != nil {
			return "*"
		}
	case OnPurchasedPaidMedia:
		if u.PurchasedPaidMedia != nil {
			return "*"
		}
	case OnCommand:
		if u.Message != nil {
			return extractCommand(u.Message)
		}
	}
	return ""
}

// extractCommand returns the command name from a message (without "/" and "@botname").
// Returns empty string if the message does not start with a bot command entity.
func extractCommand(msg *models.Message) string {
	for _, e := range msg.Entities {
		if e.Type == enums.EntityBotCommand && e.Offset == 0 {
			cmd := msg.Text[1:e.Length] // skip "/"
			if at := strings.IndexByte(cmd, '@'); at >= 0 {
				cmd = cmd[:at]
			}
			return cmd
		}
	}
	// Fallback: parse text directly if no entities (e.g. in tests).
	if strings.HasPrefix(msg.Text, "/") {
		cmd := msg.Text[1:]
		if sp := strings.IndexByte(cmd, ' '); sp >= 0 {
			cmd = cmd[:sp]
		}
		if at := strings.IndexByte(cmd, '@'); at >= 0 {
			cmd = cmd[:at]
		}
		if cmd != "" {
			return cmd
		}
	}
	return ""
}

// matchPattern checks if text matches pattern using the given match type.
func matchPattern(mt MatchType, text, pattern string) bool {
	if pattern == "" || pattern == "*" {
		return true
	}
	switch mt {
	case MatchExact:
		return text == pattern
	case MatchPrefix:
		return strings.HasPrefix(text, pattern)
	case MatchContains:
		return strings.Contains(text, pattern)
	}
	return false
}

// processUpdate finds matching handler and runs it through middleware chain.
// rawJSON is the original JSON from Telegram, used for the dev dashboard.
func (b *Bot) processUpdate(ctx context.Context, update *models.Update, rawJSON ...[]byte) {
	// Broadcast to dev dashboard if enabled.
	if b.devServer != nil {
		if len(rawJSON) > 0 && rawJSON[0] != nil {
			// Use raw JSON from Telegram — preserves all fields (e.g. callback_query.message).
			b.devServer.Broadcast(rawJSON[0])
		} else if data, err := json.Marshal(update); err == nil {
			b.devServer.Broadcast(data)
		}
	}

	ctx = withUpdate(ctx, update)
	h := b.matchHandler(update)
	if h == nil {
		return
	}

	// Build middleware chain
	final := h
	b.mu.RLock()
	middlewares := b.middlewares
	b.mu.RUnlock()

	for i := len(middlewares) - 1; i >= 0; i-- {
		next := final
		mw := middlewares[i]
		final = func(ctx context.Context, b *Bot, update *models.Update) {
			mw(ctx, b, update, next)
		}
	}

	final(ctx, b, update)
}

package gramkit

import (
	"context"

	"github.com/gram-kit/gramkit-go/models"
)

// Typed handler function types. Each receives the specific object
// instead of the full *models.Update.
type (
	MessageHandlerFunc              func(ctx context.Context, b *Bot, msg *models.Message)
	CallbackQueryHandlerFunc        func(ctx context.Context, b *Bot, cq *models.CallbackQuery)
	InlineQueryHandlerFunc          func(ctx context.Context, b *Bot, iq *models.InlineQuery)
	ChosenInlineResultHandlerFunc   func(ctx context.Context, b *Bot, cir *models.ChosenInlineResult)
	ShippingQueryHandlerFunc        func(ctx context.Context, b *Bot, sq *models.ShippingQuery)
	PreCheckoutQueryHandlerFunc     func(ctx context.Context, b *Bot, pcq *models.PreCheckoutQuery)
	PollHandlerFunc                 func(ctx context.Context, b *Bot, poll *models.Poll)
	PollAnswerHandlerFunc           func(ctx context.Context, b *Bot, pa *models.PollAnswer)
	ChatMemberUpdatedHandlerFunc    func(ctx context.Context, b *Bot, cmu *models.ChatMemberUpdated)
	ChatJoinRequestHandlerFunc      func(ctx context.Context, b *Bot, cjr *models.ChatJoinRequest)
	ChatBoostUpdatedHandlerFunc     func(ctx context.Context, b *Bot, cbu *models.ChatBoostUpdated)
	ChatBoostRemovedHandlerFunc     func(ctx context.Context, b *Bot, cbr *models.ChatBoostRemoved)
	MessageReactionHandlerFunc      func(ctx context.Context, b *Bot, mr *models.MessageReactionUpdated)
	MessageReactionCountHandlerFunc func(ctx context.Context, b *Bot, mrc *models.MessageReactionCountUpdated)
	BusinessConnectionHandlerFunc   func(ctx context.Context, b *Bot, bc *models.BusinessConnection)
	BusinessMessagesDeletedHandlerFunc func(ctx context.Context, b *Bot, bmd *models.BusinessMessagesDeleted)
	PaidMediaPurchasedHandlerFunc   func(ctx context.Context, b *Bot, pmp *models.PaidMediaPurchased)
)

// --- Message-based handlers ---

// HandleCommand registers a command handler. The command should be without the "/" prefix.
// Correctly handles /command, /command@botname, and /command args.
func (b *Bot) HandleCommand(command string, h MessageHandlerFunc) {
	b.RegisterHandler(OnCommand, command, wrapMessage(h, messageFromUpdate))
}

// HandleMessage registers a typed message handler (exact match).
func (b *Bot) HandleMessage(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnMessage, pattern, wrapMessage(h, messageFromUpdate))
}

// HandleMessageMatch registers a typed message handler with a specific match type.
func (b *Bot) HandleMessageMatch(matchType MatchType, pattern string, h MessageHandlerFunc) {
	b.RegisterHandlerMatch(OnMessage, matchType, pattern, wrapMessage(h, messageFromUpdate))
}

// HandleEditedMessage registers a typed edited message handler (exact match).
func (b *Bot) HandleEditedMessage(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnEditedMessage, pattern, wrapMessage(h, func(u *models.Update) *models.Message { return u.EditedMessage }))
}

// HandleChannelPost registers a typed channel post handler (exact match).
func (b *Bot) HandleChannelPost(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnChannelPost, pattern, wrapMessage(h, func(u *models.Update) *models.Message { return u.ChannelPost }))
}

// HandleEditedChannelPost registers a typed edited channel post handler (exact match).
func (b *Bot) HandleEditedChannelPost(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnEditedChannelPost, pattern, wrapMessage(h, func(u *models.Update) *models.Message { return u.EditedChannelPost }))
}

// HandleBusinessMessage registers a typed business message handler (exact match).
func (b *Bot) HandleBusinessMessage(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnBusinessMessage, pattern, wrapMessage(h, func(u *models.Update) *models.Message { return u.BusinessMessage }))
}

// HandleEditedBusinessMessage registers a typed edited business message handler (exact match).
func (b *Bot) HandleEditedBusinessMessage(pattern string, h MessageHandlerFunc) {
	b.RegisterHandler(OnEditedBusinessMessage, pattern, wrapMessage(h, func(u *models.Update) *models.Message { return u.EditedBusinessMessage }))
}

// --- CallbackQuery ---

// HandleCallbackQuery registers a typed callback query handler (exact match).
func (b *Bot) HandleCallbackQuery(pattern string, h CallbackQueryHandlerFunc) {
	b.RegisterHandler(OnCallbackQuery, pattern, func(ctx context.Context, b *Bot, u *models.Update) {
		if u.CallbackQuery != nil {
			h(ctx, b, u.CallbackQuery)
		}
	})
}

// HandleCallbackQueryMatch registers a typed callback query handler with a specific match type.
func (b *Bot) HandleCallbackQueryMatch(matchType MatchType, pattern string, h CallbackQueryHandlerFunc) {
	b.RegisterHandlerMatch(OnCallbackQuery, matchType, pattern, func(ctx context.Context, b *Bot, u *models.Update) {
		if u.CallbackQuery != nil {
			h(ctx, b, u.CallbackQuery)
		}
	})
}

// --- InlineQuery ---

// HandleInlineQuery registers a typed inline query handler (exact match).
func (b *Bot) HandleInlineQuery(pattern string, h InlineQueryHandlerFunc) {
	b.RegisterHandler(OnInlineQuery, pattern, func(ctx context.Context, b *Bot, u *models.Update) {
		if u.InlineQuery != nil {
			h(ctx, b, u.InlineQuery)
		}
	})
}

// HandleInlineQueryMatch registers a typed inline query handler with a specific match type.
func (b *Bot) HandleInlineQueryMatch(matchType MatchType, pattern string, h InlineQueryHandlerFunc) {
	b.RegisterHandlerMatch(OnInlineQuery, matchType, pattern, func(ctx context.Context, b *Bot, u *models.Update) {
		if u.InlineQuery != nil {
			h(ctx, b, u.InlineQuery)
		}
	})
}

// --- ChosenInlineResult ---

// HandleChosenInlineResult registers a typed chosen inline result handler.
func (b *Bot) HandleChosenInlineResult(pattern string, h ChosenInlineResultHandlerFunc) {
	b.RegisterHandler(OnChosenInlineResult, pattern, func(ctx context.Context, b *Bot, u *models.Update) {
		if u.ChosenInlineResult != nil {
			h(ctx, b, u.ChosenInlineResult)
		}
	})
}

// --- ShippingQuery ---

// HandleShippingQuery registers a typed shipping query handler.
func (b *Bot) HandleShippingQuery(h ShippingQueryHandlerFunc) {
	b.RegisterHandlerMatch(OnShippingQuery, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.ShippingQuery != nil {
			h(ctx, b, u.ShippingQuery)
		}
	})
}

// --- PreCheckoutQuery ---

// HandlePreCheckoutQuery registers a typed pre-checkout query handler.
func (b *Bot) HandlePreCheckoutQuery(h PreCheckoutQueryHandlerFunc) {
	b.RegisterHandlerMatch(OnPreCheckoutQuery, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.PreCheckoutQuery != nil {
			h(ctx, b, u.PreCheckoutQuery)
		}
	})
}

// --- Poll ---

// HandlePoll registers a typed poll handler.
func (b *Bot) HandlePoll(h PollHandlerFunc) {
	b.RegisterHandlerMatch(OnPoll, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.Poll != nil {
			h(ctx, b, u.Poll)
		}
	})
}

// --- PollAnswer ---

// HandlePollAnswer registers a typed poll answer handler.
func (b *Bot) HandlePollAnswer(h PollAnswerHandlerFunc) {
	b.RegisterHandlerMatch(OnPollAnswer, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.PollAnswer != nil {
			h(ctx, b, u.PollAnswer)
		}
	})
}

// --- ChatMemberUpdated ---

// HandleMyChatMember registers a typed handler for changes to the bot's chat member status.
func (b *Bot) HandleMyChatMember(h ChatMemberUpdatedHandlerFunc) {
	b.RegisterHandlerMatch(OnMyChatMember, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.MyChatMember != nil {
			h(ctx, b, u.MyChatMember)
		}
	})
}

// HandleChatMember registers a typed handler for changes to a chat member's status.
func (b *Bot) HandleChatMember(h ChatMemberUpdatedHandlerFunc) {
	b.RegisterHandlerMatch(OnChatMember, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.ChatMember != nil {
			h(ctx, b, u.ChatMember)
		}
	})
}

// --- ChatJoinRequest ---

// HandleChatJoinRequest registers a typed chat join request handler.
func (b *Bot) HandleChatJoinRequest(h ChatJoinRequestHandlerFunc) {
	b.RegisterHandlerMatch(OnChatJoinRequest, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.ChatJoinRequest != nil {
			h(ctx, b, u.ChatJoinRequest)
		}
	})
}

// --- ChatBoost ---

// HandleChatBoost registers a typed chat boost handler.
func (b *Bot) HandleChatBoost(h ChatBoostUpdatedHandlerFunc) {
	b.RegisterHandlerMatch(OnChatBoost, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.ChatBoost != nil {
			h(ctx, b, u.ChatBoost)
		}
	})
}

// HandleRemovedChatBoost registers a typed removed chat boost handler.
func (b *Bot) HandleRemovedChatBoost(h ChatBoostRemovedHandlerFunc) {
	b.RegisterHandlerMatch(OnRemovedChatBoost, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.RemovedChatBoost != nil {
			h(ctx, b, u.RemovedChatBoost)
		}
	})
}

// --- MessageReaction ---

// HandleMessageReaction registers a typed message reaction handler.
func (b *Bot) HandleMessageReaction(h MessageReactionHandlerFunc) {
	b.RegisterHandlerMatch(OnMessageReaction, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.MessageReaction != nil {
			h(ctx, b, u.MessageReaction)
		}
	})
}

// HandleMessageReactionCount registers a typed message reaction count handler.
func (b *Bot) HandleMessageReactionCount(h MessageReactionCountHandlerFunc) {
	b.RegisterHandlerMatch(OnMessageReactionCount, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.MessageReactionCount != nil {
			h(ctx, b, u.MessageReactionCount)
		}
	})
}

// --- Business ---

// HandleBusinessConnection registers a typed business connection handler.
func (b *Bot) HandleBusinessConnection(h BusinessConnectionHandlerFunc) {
	b.RegisterHandlerMatch(OnBusinessConnection, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.BusinessConnection != nil {
			h(ctx, b, u.BusinessConnection)
		}
	})
}

// HandleDeletedBusinessMessages registers a typed deleted business messages handler.
func (b *Bot) HandleDeletedBusinessMessages(h BusinessMessagesDeletedHandlerFunc) {
	b.RegisterHandlerMatch(OnDeletedBusinessMessages, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.DeletedBusinessMessages != nil {
			h(ctx, b, u.DeletedBusinessMessages)
		}
	})
}

// --- PaidMedia ---

// HandlePurchasedPaidMedia registers a typed purchased paid media handler.
func (b *Bot) HandlePurchasedPaidMedia(h PaidMediaPurchasedHandlerFunc) {
	b.RegisterHandlerMatch(OnPurchasedPaidMedia, MatchExact, "*", func(ctx context.Context, b *Bot, u *models.Update) {
		if u.PurchasedPaidMedia != nil {
			h(ctx, b, u.PurchasedPaidMedia)
		}
	})
}

// --- internal helpers ---

func messageFromUpdate(u *models.Update) *models.Message { return u.Message }

func wrapMessage(h MessageHandlerFunc, extract func(*models.Update) *models.Message) HandlerFunc {
	return func(ctx context.Context, b *Bot, update *models.Update) {
		if msg := extract(update); msg != nil {
			h(ctx, b, msg)
		}
	}
}

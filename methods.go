package gramkit

import (
	"context"

	"github.com/gram-kit/gramkit-go/internal/network"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

// --- Getting Updates --- //

// GetUpdates receives incoming updates using long polling. Returns an Array of Update objects.
func (b *Bot) GetUpdates(ctx context.Context, params ...params.GetUpdates) ([]models.Update, error) {
	var p any
	if len(params) > 0 {
		p = params[0]
	}
	return network.UseMethod[[]models.Update](b.client, ctx, "getUpdates", p)
}

// SetWebhook specifies a URL to receive incoming updates via an outgoing webhook. Returns True on success.
func (b *Bot) SetWebhook(ctx context.Context, params params.SetWebhook) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setWebhook", params)
}

// DeleteWebhook removes webhook integration. Returns True on success.
func (b *Bot) DeleteWebhook(ctx context.Context, params ...params.DeleteWebhook) (bool, error) {
	var p any
	if len(params) > 0 {
		p = params[0]
	}
	return network.UseMethod[bool](b.client, ctx, "deleteWebhook", p)
}

// GetWebhookInfo gets current webhook status. Returns a WebhookInfo object.
func (b *Bot) GetWebhookInfo(ctx context.Context) (models.WebhookInfo, error) {
	return network.UseMethod[models.WebhookInfo](b.client, ctx, "getWebhookInfo", nil)
}

// --- Available Methods --- //

// GetMe returns basic information about the bot as a User object.
func (b *Bot) GetMe(ctx context.Context) (models.User, error) {
	return network.UseMethod[models.User](b.client, ctx, "getMe", nil)
}

// LogOut logs out from the cloud Bot API server. Returns True on success.
func (b *Bot) LogOut(ctx context.Context) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "logOut", nil)
}

// Close closes the bot instance before moving it from one local server to another. Returns True on success.
func (b *Bot) Close(ctx context.Context) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "close", nil)
}

// SendMessage sends a text message. On success, the sent Message is returned.
func (b *Bot) SendMessage(ctx context.Context, params params.SendMessage) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendMessage", params)
}

// ForwardMessage forwards a message of any kind. On success, the sent Message is returned.
func (b *Bot) ForwardMessage(ctx context.Context, params params.ForwardMessage) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "forwardMessage", params)
}

// ForwardMessages forwards multiple messages of any kind. On success, an array of MessageID is returned.
func (b *Bot) ForwardMessages(ctx context.Context, params params.ForwardMessages) ([]models.MessageID, error) {
	return network.UseMethod[[]models.MessageID](b.client, ctx, "forwardMessages", params)
}

// CopyMessage copies a message. Returns the MessageID of the sent message on success.
func (b *Bot) CopyMessage(ctx context.Context, params params.CopyMessage) (models.MessageID, error) {
	return network.UseMethod[models.MessageID](b.client, ctx, "copyMessage", params)
}

// CopyMessages copies multiple messages. On success, an array of MessageID is returned.
func (b *Bot) CopyMessages(ctx context.Context, params params.CopyMessages) ([]models.MessageID, error) {
	return network.UseMethod[[]models.MessageID](b.client, ctx, "copyMessages", params)
}

// SendPhoto sends a photo. On success, the sent Message is returned.
func (b *Bot) SendPhoto(ctx context.Context, params params.SendPhoto) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendPhoto", params)
}

// SendAudio sends an audio file. On success, the sent Message is returned.
func (b *Bot) SendAudio(ctx context.Context, params params.SendAudio) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendAudio", params)
}

// SendDocument sends a general file. On success, the sent Message is returned.
func (b *Bot) SendDocument(ctx context.Context, params params.SendDocument) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendDocument", params)
}

// SendVideo sends a video file. On success, the sent Message is returned.
func (b *Bot) SendVideo(ctx context.Context, params params.SendVideo) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendVideo", params)
}

// SendAnimation sends an animation (GIF or video without sound). On success, the sent Message is returned.
func (b *Bot) SendAnimation(ctx context.Context, params params.SendAnimation) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendAnimation", params)
}

// SendVoice sends a voice message. On success, the sent Message is returned.
func (b *Bot) SendVoice(ctx context.Context, params params.SendVoice) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendVoice", params)
}

// SendVideoNote sends a video note (rounded square video). On success, the sent Message is returned.
func (b *Bot) SendVideoNote(ctx context.Context, params params.SendVideoNote) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendVideoNote", params)
}

// SendPaidMedia sends paid media. On success, the sent Message is returned.
func (b *Bot) SendPaidMedia(ctx context.Context, params params.SendPaidMedia) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendPaidMedia", params)
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album. On success, an array of Messages is returned.
func (b *Bot) SendMediaGroup(ctx context.Context, params params.SendMediaGroup) ([]models.Message, error) {
	return network.UseMethod[[]models.Message](b.client, ctx, "sendMediaGroup", params)
}

// SendLocation sends a point on the map. On success, the sent Message is returned.
func (b *Bot) SendLocation(ctx context.Context, params params.SendLocation) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendLocation", params)
}

// SendVenue sends information about a venue. On success, the sent Message is returned.
func (b *Bot) SendVenue(ctx context.Context, params params.SendVenue) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendVenue", params)
}

// SendContact sends a phone contact. On success, the sent Message is returned.
func (b *Bot) SendContact(ctx context.Context, params params.SendContact) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendContact", params)
}

// SendChecklist sends a checklist. On success, the sent Message is returned.
func (b *Bot) SendChecklist(ctx context.Context, params params.SendChecklist) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendChecklist", params)
}

// SendPoll sends a native poll. On success, the sent Message is returned.
func (b *Bot) SendPoll(ctx context.Context, params params.SendPoll) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendPoll", params)
}

// SendDice sends an animated emoji with a random value. On success, the sent Message is returned.
func (b *Bot) SendDice(ctx context.Context, params params.SendDice) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendDice", params)
}

// SendMessageDraft streams a partial message draft. Returns True on success.
func (b *Bot) SendMessageDraft(ctx context.Context, params params.SendMessageDraft) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "sendMessageDraft", params)
}

// SendChatAction sends a chat action (typing, uploading, etc). Returns True on success.
func (b *Bot) SendChatAction(ctx context.Context, params params.SendChatAction) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "sendChatAction", params)
}

// SetMessageReaction changes the chosen reactions on a message. Returns True on success.
func (b *Bot) SetMessageReaction(ctx context.Context, params params.SetMessageReaction) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMessageReaction", params)
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a models.UserProfilePhotos object.
func (b *Bot) GetUserProfilePhotos(ctx context.Context, params params.GetUserProfilePhotos) (models.UserProfilePhotos, error) {
	return network.UseMethod[models.UserProfilePhotos](b.client, ctx, "getUserProfilePhotos", params)
}

// GetUserProfileAudios Use this method to get a list of profile audios for a user. Returns a models.UserProfileAudios object.
func (b *Bot) GetUserProfileAudios(ctx context.Context, params params.GetUserProfileAudios) (models.UserProfileAudios, error) {
	return network.UseMethod[models.UserProfileAudios](b.client, ctx, "getUserProfileAudios", params)
}

// SetUserEmojiStatus Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method RequestEmojiStatusAccess. Returns True on success.
func (b *Bot) SetUserEmojiStatus(ctx context.Context, params params.SetUserEmojiStatus) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setUserEmojiStatus", params)
}

// GetFile Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a models.File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling GetFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (b *Bot) GetFile(ctx context.Context, params params.GetFile) (models.File, error) {
	return network.UseMethod[models.File](b.client, ctx, "getFile", params)
}

// BanChatMember Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatMember(ctx context.Context, params params.BanChatMember) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "banChatMember", params)
}

// UnbanChatMember Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (b *Bot) UnbanChatMember(ctx context.Context, params params.UnbanChatMember) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unbanChatMember", params)
}

// RestrictChatMember Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(ctx context.Context, params params.RestrictChatMember) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "restrictChatMember", params)
}

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(ctx context.Context, params params.PromoteChatMember) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "promoteChatMember", params)
}

// SetChatAdministratorCustomTitle Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (b *Bot) SetChatAdministratorCustomTitle(ctx context.Context, params params.SetChatAdministratorCustomTitle) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatAdministratorCustomTitle", params)
}

// SetChatMemberTag Use this method to set a tag for a regular member in a group or a supergroup. The bot must be an administrator in the chat for this to work and must have the can_manage_tags administrator right. Returns True on success.
func (b *Bot) SetChatMemberTag(ctx context.Context, params params.SetChatMemberTag) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatMemberTag", params)
}

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatSenderChat(ctx context.Context, params params.BanChatSenderChat) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "banChatSenderChat", params)
}

// UnbanChatSenderChat Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) UnbanChatSenderChat(ctx context.Context, params params.UnbanChatSenderChat) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unbanChatSenderChat", params)
}

// SetChatPermissions Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (b *Bot) SetChatPermissions(ctx context.Context, params params.SetChatPermissions) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatPermissions", params)
}

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using ExportChatInviteLink or by calling the GetChat method. If your bot needs to generate a new primary invite link replacing its previous one, use ExportChatInviteLink again.
func (b *Bot) ExportChatInviteLink(ctx context.Context, params params.ExportChatInviteLink) (string, error) {
	return network.UseMethod[string](b.client, ctx, "exportChatInviteLink", params)
}

// CreateChatInviteLink Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method RevokeChatInviteLink. Returns the new invite link as models.ChatInviteLink object.
func (b *Bot) CreateChatInviteLink(ctx context.Context, params params.CreateChatInviteLink) (models.ChatInviteLink, error) {
	return network.UseMethod[models.ChatInviteLink](b.client, ctx, "createChatInviteLink", params)
}

// EditChatInviteLink Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a models.ChatInviteLink object.
func (b *Bot) EditChatInviteLink(ctx context.Context, params params.EditChatInviteLink) (models.ChatInviteLink, error) {
	return network.UseMethod[models.ChatInviteLink](b.client, ctx, "editChatInviteLink", params)
}

// CreateChatSubscriptionInviteLink Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method EditChatSubscriptionInviteLink or revoked using the method RevokeChatInviteLink. Returns the new invite link as a models.ChatInviteLink object.
func (b *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, params params.CreateChatSubscriptionInviteLink) (models.ChatInviteLink, error) {
	return network.UseMethod[models.ChatInviteLink](b.client, ctx, "createChatSubscriptionInviteLink", params)
}

// EditChatSubscriptionInviteLink Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a models.ChatInviteLink object.
func (b *Bot) EditChatSubscriptionInviteLink(ctx context.Context, params params.EditChatSubscriptionInviteLink) (models.ChatInviteLink, error) {
	return network.UseMethod[models.ChatInviteLink](b.client, ctx, "editChatSubscriptionInviteLink", params)
}

// RevokeChatInviteLink Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as models.ChatInviteLink object.
func (b *Bot) RevokeChatInviteLink(ctx context.Context, params params.RevokeChatInviteLink) (models.ChatInviteLink, error) {
	return network.UseMethod[models.ChatInviteLink](b.client, ctx, "revokeChatInviteLink", params)
}

// ApproveChatJoinRequest Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) ApproveChatJoinRequest(ctx context.Context, params params.ApproveChatJoinRequest) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "approveChatJoinRequest", params)
}

// DeclineChatJoinRequest Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) DeclineChatJoinRequest(ctx context.Context, params params.DeclineChatJoinRequest) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "declineChatJoinRequest", params)
}

// SetChatPhoto Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatPhoto(ctx context.Context, params params.SetChatPhoto) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatPhoto", params)
}

// DeleteChatPhoto Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) DeleteChatPhoto(ctx context.Context, params params.DeleteChatPhoto) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteChatPhoto", params)
}

// SetChatTitle Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatTitle(ctx context.Context, params params.SetChatTitle) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatTitle", params)
}

// SetChatDescription Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatDescription(ctx context.Context, params params.SetChatDescription) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatDescription", params)
}

// PinChatMessage Use this method to add a message to the list of pinned messages in a chat. In private chats and channel direct messages chats, all non-service messages can be pinned. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to pin messages in groups and channels respectively. Returns True on success.
func (b *Bot) PinChatMessage(ctx context.Context, params params.PinChatMessage) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "pinChatMessage", params)
}

// UnpinChatMessage Use this method to remove a message from the list of pinned messages in a chat. Returns True on success.
func (b *Bot) UnpinChatMessage(ctx context.Context, params params.UnpinChatMessage) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unpinChatMessage", params)
}

// UnpinAllChatMessages Use this method to clear the list of pinned messages in a chat. Returns True on success.
func (b *Bot) UnpinAllChatMessages(ctx context.Context, params params.UnpinAllChatMessage) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unpinAllChatMessages", params)
}

// LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(ctx context.Context, params params.LeaveChat) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "leaveChat", params)
}

// GetChat Use this method to get up-to-date information about the chat. Returns a models.ChatFullInfo object on success.
func (b *Bot) GetChat(ctx context.Context, params params.GetChat) (models.ChatFullInfo, error) {
	return network.UseMethod[models.ChatFullInfo](b.client, ctx, "getChat", params)
}

// GetChatAdministrators Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of models.ChatMember objects.
func (b *Bot) GetChatAdministrators(ctx context.Context, params params.GetChatAdministrators) ([]models.ChatMember, error) {
	return network.UseMethod[[]models.ChatMember](b.client, ctx, "getChatAdministrators", params)
}

// GetChatMemberCount Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMemberCount(ctx context.Context, params params.GetChatMemberCount) (int, error) {
	return network.UseMethod[int](b.client, ctx, "getChatMemberCount", params)
}

// GetChatMember Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a models.ChatMember object on success.
func (b *Bot) GetChatMember(ctx context.Context, params params.GetChatMember) (models.ChatMember, error) {
	return network.UseMethod[models.ChatMember](b.client, ctx, "getChatMember", params)
}

// SetChatStickerSet Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in GetChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) SetChatStickerSet(ctx context.Context, params params.SetChatStickerSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatStickerSet", params)
}

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in GetChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) DeleteChatStickerSet(ctx context.Context, params params.DeleteChatStickerSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteChatStickerSet", params)
}

// GetForumTopicIconStickers Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of models.Sticker objects.
func (b *Bot) GetForumTopicIconStickers(ctx context.Context) ([]models.Sticker, error) {
	return network.UseMethod[[]models.Sticker](b.client, ctx, "getForumTopicIconStickers", nil)
}

// CreateForumTopic Use this method to create a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator right. Returns information about the created topic as a models.ForumTopic object.
func (b *Bot) CreateForumTopic(ctx context.Context, params params.CreateForumTopic) (models.ForumTopic, error) {
	return network.UseMethod[models.ForumTopic](b.client, ctx, "createForumTopic", params)
}

// EditForumTopic Use this method to edit name and icon of a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) EditForumTopic(ctx context.Context, params params.EditForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editForumTopic", params)
}

// CloseForumTopic Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) CloseForumTopic(ctx context.Context, params params.CloseForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "closeForumTopic", params)
}

// ReopenForumTopic Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) ReopenForumTopic(ctx context.Context, params params.ReopenForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "reopenForumTopic", params)
}

// DeleteForumTopic Use this method to delete a forum topic along with all its messages in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (b *Bot) DeleteForumTopic(ctx context.Context, params params.DeleteForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteForumTopic", params)
}

// UnpinAllForumTopicMessages Use this method to clear the list of pinned messages in a forum topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllForumTopicMessages(ctx context.Context, params params.UnpinAllForumTopicMessages) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unpinAllForumTopicMessages", params)
}

// EditGeneralForumTopic Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) EditGeneralForumTopic(ctx context.Context, params params.EditGeneralForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editGeneralForumTopic", params)
}

// CloseGeneralForumTopic Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) CloseGeneralForumTopic(ctx context.Context, params params.CloseGeneralForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "closeGeneralForumTopic", params)
}

// ReopenGeneralForumTopic Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (b *Bot) ReopenGeneralForumTopic(ctx context.Context, params params.ReopenGeneralForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "reopenGeneralForumTopic", params)
}

// HideGeneralForumTopic Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (b *Bot) HideGeneralForumTopic(ctx context.Context, params params.HideGeneralForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "hideGeneralForumTopic", params)
}

// UnhideGeneralForumTopic Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) UnhideGeneralForumTopic(ctx context.Context, params params.UnhideGeneralForumTopic) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unhideGeneralForumTopic", params)
}

// UnpinAllGeneralForumTopicMessages Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, params params.UnpinAllGeneralForumTopicMessages) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "unpinAllGeneralForumTopicMessages", params)
}

// AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
func (b *Bot) AnswerCallbackQuery(ctx context.Context, params params.AnswerCallbackQuery) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "answerCallbackQuery", params)
}

// GetUserChatBoosts Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a models.UserChatBoosts object.
func (b *Bot) GetUserChatBoosts(ctx context.Context, params params.GetUserChatBoosts) (models.UserChatBoosts, error) {
	return network.UseMethod[models.UserChatBoosts](b.client, ctx, "getUserChatBoosts", params)
}

// GetBusinessConnection Use this method to get information about the connection of the bot with a business account. Returns a models.BusinessConnection object on success.
func (b *Bot) GetBusinessConnection(ctx context.Context, params params.GetBusinessConnection) (models.BusinessConnection, error) {
	return network.UseMethod[models.BusinessConnection](b.client, ctx, "getBusinessConnection", params)
}

// SetMyCommands Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (b *Bot) SetMyCommands(ctx context.Context, params params.SetMyCommands) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyCommands", params)
}

// DeleteMyCommands Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (b *Bot) DeleteMyCommands(ctx context.Context, params params.DeleteMyCommands) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteMyCommands", params)
}

// GetMyCommands Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of models.BotCommand objects. If commands aren't set, an empty list is returned.
func (b *Bot) GetMyCommands(ctx context.Context, params params.GetMyCommands) ([]models.BotCommand, error) {
	return network.UseMethod[[]models.BotCommand](b.client, ctx, "getMyCommands", params)
}

// SetMyName Use this method to change the bot's name. Returns True on success.
func (b *Bot) SetMyName(ctx context.Context, params params.SetMyName) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyName", params)
}

// GetMyName Use this method to get the current bot name for the given user language. Returns models.BotName on success.
func (b *Bot) GetMyName(ctx context.Context, params params.GetMyName) (models.BotName, error) {
	return network.UseMethod[models.BotName](b.client, ctx, "getMyName", params)
}

// SetMyDescription Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
func (b *Bot) SetMyDescription(ctx context.Context, params params.SetMyDescription) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyDescription", params)
}

// GetMyDescription Use this method to get the current bot description for the given user language. Returns models.BotDescription on success.
func (b *Bot) GetMyDescription(ctx context.Context, params params.GetMyDescription) (models.BotDescription, error) {
	return network.UseMethod[models.BotDescription](b.client, ctx, "getMyDescription", params)
}

// SetMyShortDescription Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
func (b *Bot) SetMyShortDescription(ctx context.Context, params params.SetMyShortDescription) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyShortDescription", params)
}

// GetMyShortDescription Use this method to get the current bot short description for the given user language. Returns models.BotShortDescription on success.
func (b *Bot) GetMyShortDescription(ctx context.Context, params params.GetMyShortDescription) (models.BotShortDescription, error) {
	return network.UseMethod[models.BotShortDescription](b.client, ctx, "getMyShortDescription", params)
}

// SetMyProfilePhoto Changes the profile photo of the bot. Returns True on success.
func (b *Bot) SetMyProfilePhoto(ctx context.Context, params params.SetMyProfilePhoto) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyProfilePhoto", params)
}

// RemoveMyProfilePhoto Removes the profile photo of the bot. Requires no parameters. Returns True on success.
func (b *Bot) RemoveMyProfilePhoto(ctx context.Context) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "removeMyProfilePhoto", nil)
}

// SetChatMenuButton Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (b *Bot) SetChatMenuButton(ctx context.Context, params params.SetChatMenuButton) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setChatMenuButton", params)
}

// GetChatMenuButton Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns models.MenuButton on success.
func (b *Bot) GetChatMenuButton(ctx context.Context, params params.GetChatMenuButton) (models.MenuButton, error) {
	return network.UseMethod[models.MenuButton](b.client, ctx, "getChatMenuButton", params)
}

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
func (b *Bot) SetMyDefaultAdministratorRights(ctx context.Context, params params.SetMyDefaultAdministratorRights) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setMyDefaultAdministratorRights", params)
}

// GetMyDefaultAdministratorRights Use this method to get the current default administrator rights of the bot. Returns models.ChatAdministratorRights on success.
func (b *Bot) GetMyDefaultAdministratorRights(ctx context.Context, params params.GetMyDefaultAdministratorRights) (models.ChatAdministratorRights, error) {
	return network.UseMethod[models.ChatAdministratorRights](b.client, ctx, "getMyDefaultAdministratorRights", params)
}

// GetAvailableGifts Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a models.Gifts object.
func (b *Bot) GetAvailableGifts(ctx context.Context) (models.Gifts, error) {
	return network.UseMethod[models.Gifts](b.client, ctx, "getAvailableGifts", nil)
}

// SendGift Sends a gift to the given user or channel chat. The gift can't be converted to Telegram Stars by the receiver. Returns True on success.
func (b *Bot) SendGift(ctx context.Context, params params.SendGift) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "sendGift", params)
}

// GiftPremiumSubscription Gifts a Telegram Premium subscription to the given user. Returns True on success.
func (b *Bot) GiftPremiumSubscription(ctx context.Context, params params.GiftPremiumSubscription) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "giftPremiumSubscription", params)
}

// VerifyUser Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
func (b *Bot) VerifyUser(ctx context.Context, params params.VerifyUser) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "verifyUser", params)
}

// VerifyChat Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
func (b *Bot) VerifyChat(ctx context.Context, params params.VerifyChat) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "verifyChat", params)
}

// RemoveUserVerification Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
func (b *Bot) RemoveUserVerification(ctx context.Context, params params.RemoveUserVerification) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "removeUserVerification", params)
}

// RemoveChatVerification Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
func (b *Bot) RemoveChatVerification(ctx context.Context, params params.RemoveChatVerification) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "removeChatVerification", params)
}

// ReadBusinessMessage Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.
func (b *Bot) ReadBusinessMessage(ctx context.Context, params params.ReadBusinessMessage) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "readBusinessMessage", params)
}

// DeleteBusinessMessages Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.
func (b *Bot) DeleteBusinessMessages(ctx context.Context, params params.DeleteBusinessMessages) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteBusinessMessages", params)
}

// SetBusinessAccountName Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.
func (b *Bot) SetBusinessAccountName(ctx context.Context, params params.SetBusinessAccountName) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setBusinessAccountName", params)
}

// SetBusinessAccountUsername Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.
func (b *Bot) SetBusinessAccountUsername(ctx context.Context, params params.SetBusinessAccountUsername) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setBusinessAccountUsername", params)
}

// SetBusinessAccountBio Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.
func (b *Bot) SetBusinessAccountBio(ctx context.Context, params params.SetBusinessAccountBio) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setBusinessAccountBio", params)
}

// SetBusinessAccountProfilePhoto Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
func (b *Bot) SetBusinessAccountProfilePhoto(ctx context.Context, params params.SetBusinessAccountProfilePhoto) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setBusinessAccountProfilePhoto", params)
}

// RemoveBusinessAccountProfilePhoto Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
func (b *Bot) RemoveBusinessAccountProfilePhoto(ctx context.Context, params params.RemoveBusinessAccountProfilePhoto) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "removeBusinessAccountProfilePhoto", params)
}

// SetBusinessAccountGiftSettings Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.
func (b *Bot) SetBusinessAccountGiftSettings(ctx context.Context, params params.SetBusinessAccountGiftSettings) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setBusinessAccountGiftSettings", params)
}

// GetBusinessAccountStarBalance Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns models.StarAmount on success.
func (b *Bot) GetBusinessAccountStarBalance(ctx context.Context, params params.GetBusinessAccountStarBalance) (models.StarAmount, error) {
	return network.UseMethod[models.StarAmount](b.client, ctx, "getBusinessAccountStarBalance", params)
}

// TransferBusinessAccountStars Transfers Telegram Stars from the business account balance to the bot's balance. Requires the can_transfer_stars business bot right. Returns True on success.
func (b *Bot) TransferBusinessAccountStars(ctx context.Context, params params.TransferBusinessAccountStars) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "transferBusinessAccountStars", params)
}

// GetBusinessAccountGifts Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns models.OwnedGifts on success.
func (b *Bot) GetBusinessAccountGifts(ctx context.Context, params params.GetBusinessAccountGifts) (models.OwnedGifts, error) {
	return network.UseMethod[models.OwnedGifts](b.client, ctx, "getBusinessAccountGifts", params)
}

// GetUserGifts Returns the gifts owned and hosted by a user. Returns models.OwnedGifts on success.
func (b *Bot) GetUserGifts(ctx context.Context, params params.GetUserGifts) (models.OwnedGifts, error) {
	return network.UseMethod[models.OwnedGifts](b.client, ctx, "getUserGifts", params)
}

// GetChatGifts Returns the gifts owned by a chat. Returns models.OwnedGifts on success.
func (b *Bot) GetChatGifts(ctx context.Context, params params.GetChatGifts) (models.OwnedGifts, error) {
	return network.UseMethod[models.OwnedGifts](b.client, ctx, "getChatGifts", params)
}

// ConvertGiftToStars Converts a given regular gift to Telegram Stars. Requires the can_convert_gifts_to_stars business bot right. Returns True on success.
func (b *Bot) ConvertGiftToStars(ctx context.Context, params params.ConvertGiftToStars) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "convertGiftToStars", params)
}

// UpgradeGift Upgrades a given regular gift to a unique gift. Requires the can_transfer_and_upgrade_gifts business bot right. Additionally requires the can_transfer_stars business bot right if the upgrade is paid. Returns True on success.
func (b *Bot) UpgradeGift(ctx context.Context, params params.UpgradeGift) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "upgradeGift", params)
}

// TransferGift Transfers an owned unique gift to another user. Requires the can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars business bot right if the transfer is paid. Returns True on success.
func (b *Bot) TransferGift(ctx context.Context, params params.TransferGift) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "transferGift", params)
}

// PostStory Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns models.Story on success.
func (b *Bot) PostStory(ctx context.Context, params params.PostStory) (models.Story, error) {
	return network.UseMethod[models.Story](b.client, ctx, "postStory", params)
}

// RepostStory Reposts a story on behalf of a business account from another business account. Both business accounts must be managed by the same bot, and the story on the source account must have been posted (or reposted) by the bot. Requires the can_manage_stories business bot right for both business accounts. Returns models.Story on success.
func (b *Bot) RepostStory(ctx context.Context, params params.RepostStory) (models.Story, error) {
	return network.UseMethod[models.Story](b.client, ctx, "repostStory", params)
}

// EditStory Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns models.Story on success.
func (b *Bot) EditStory(ctx context.Context, params params.EditStory) (models.Story, error) {
	return network.UseMethod[models.Story](b.client, ctx, "editStory", params)
}

// DeleteStory Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.
func (b *Bot) DeleteStory(ctx context.Context, params params.DeleteStory) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteStory", params)
}

// --- Updating messages --- //

// EditMessageText Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited models.Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (b *Bot) EditMessageText(ctx context.Context, params params.EditMessageText) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editMessageText", params)
}

// EditMessageCaption Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited models.Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (b *Bot) EditMessageCaption(ctx context.Context, params params.EditMessageCaption) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editMessageCaption", params)
}

// EditMessageMedia Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited models.Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (b *Bot) EditMessageMedia(ctx context.Context, params params.EditMessageMedia) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editMessageMedia", params)
}

// EditMessageLiveLocation Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to StopMessageLiveLocation. On success, if the edited message is not an inline message, the edited models.Message is returned, otherwise True is returned.
func (b *Bot) EditMessageLiveLocation(ctx context.Context, params params.EditMessageLiveLocation) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editMessageLiveLocation", params)
}

// StopMessageLiveLocation Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited models.Message is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(ctx context.Context, params params.StopMessageLiveLocation) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "stopMessageLiveLocation", params)
}

// EditMessageChecklist Use this method to edit a checklist on behalf of a connected business account. On success, the edited models.Message is returned.
func (b *Bot) EditMessageChecklist(ctx context.Context, params params.EditMessageChecklist) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "editMessageChecklist", params)
}

// EditMessageReplyMarkup Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited models.Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (b *Bot) EditMessageReplyMarkup(ctx context.Context, params params.EditMessageReplyMarkup) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editMessageReplyMarkup", params)
}

// StopPoll Use this method to stop a poll which was sent by the bot. On success, the stopped models.Poll is returned.
func (b *Bot) StopPoll(ctx context.Context, params params.StopPoll) (models.Poll, error) {
	return network.UseMethod[models.Poll](b.client, ctx, "stopPoll", params)
}

// ApproveSuggestedPost Use this method to approve a suggested post in a direct messages chat. The bot must have the 'can_post_messages' administrator right in the corresponding channel chat. Returns True on success.
func (b *Bot) ApproveSuggestedPost(ctx context.Context, params params.ApproveSuggestedPost) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "approveSuggestedPost", params)
}

// DeclineSuggestedPost Use this method to decline a suggested post in a direct messages chat. The bot must have the 'can_manage_direct_messages' administrator right in the corresponding channel chat. Returns True on success.
func (b *Bot) DeclineSuggestedPost(ctx context.Context, params params.DeclineSuggestedPost) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "declineSuggestedPost", params)
}

// DeleteMessage Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages administrator right in a supergroup or a channel, it can delete any message there.
// - If the bot has can_manage_direct_messages administrator right in a channel, it can delete any message in the corresponding direct messages chat.
// Returns True on success.
func (b *Bot) DeleteMessage(ctx context.Context, params params.DeleteMessage) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteMessage", params)
}

// DeleteMessages Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
func (b *Bot) DeleteMessages(ctx context.Context, params params.DeleteMessages) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteMessages", params)
}

// --- Stickers --- //

// SendSticker Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent models.Message is returned.
func (b *Bot) SendSticker(ctx context.Context, params params.SendSticker) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendSticker", params)
}

// GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
func (b *Bot) GetStickerSet(ctx context.Context, params params.GetStickerSet) (models.StickerSet, error) {
	return network.UseMethod[models.StickerSet](b.client, ctx, "getStickerSet", params)
}

// GetCustomEmojiStickers Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (b *Bot) GetCustomEmojiStickers(ctx context.Context, params params.GetCustomEmojiStickers) ([]models.Sticker, error) {
	return network.UseMethod[[]models.Sticker](b.client, ctx, "getCustomEmojiStickers", params)
}

// UploadStickerFile Use this method to upload a file with a sticker for later use in the CreateNewStickerSet, AddStickerToSet, or ReplaceStickerInSet methods (the file can be used multiple times). Returns the uploaded models.File on success.
func (b *Bot) UploadStickerFile(ctx context.Context, params params.UploadStickerFile) (models.File, error) {
	return network.UseMethod[models.File](b.client, ctx, "uploadStickerFile", params)
}

// CreateNewStickerSet Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
func (b *Bot) CreateNewStickerSet(ctx context.Context, params params.CreateNewStickerSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "createNewStickerSet", params)
}

// AddStickerToSet Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
func (b *Bot) AddStickerToSet(ctx context.Context, params params.AddStickerToSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "addStickerToSet", params)
}

// SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(ctx context.Context, params params.SetStickerPositionInSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerPositionInSet", params)
}

// DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (b *Bot) DeleteStickerFromSet(ctx context.Context, params params.DeleteStickerFromSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteStickerFromSet", params)
}

// ReplaceStickerInSet Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling DeleteStickerFromSet, then AddStickerToSet, then SetStickerPositionInSet. Returns True on success.
func (b *Bot) ReplaceStickerInSet(ctx context.Context, params params.ReplaceStickerInSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "replaceStickerInSet", params)
}

// SetStickerEmojiList Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerEmojiList(ctx context.Context, params params.SetStickerEmojiList) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerEmojiList", params)
}

// SetStickerKeywords Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerKeywords(ctx context.Context, params params.SetStickerKeywords) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerKeywords", params)
}

// SetStickerMaskPosition Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
func (b *Bot) SetStickerMaskPosition(ctx context.Context, params params.SetStickerMaskPosition) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerMaskPosition", params)
}

// SetStickerSetTitle Use this method to set the title of a created sticker set. Returns True on success.
func (b *Bot) SetStickerSetTitle(ctx context.Context, params params.SetStickerSetTitle) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerSetTitle", params)
}

// SetStickerSetThumbnail Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
func (b *Bot) SetStickerSetThumbnail(ctx context.Context, params params.SetStickerSetThumbnail) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setStickerSetThumbnail", params)
}

// SetCustomEmojiStickerSetThumbnail Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
func (b *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, params params.SetCustomEmojiStickerSetThumbnail) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setCustomEmojiStickerSetThumbnail", params)
}

// DeleteStickerSet Use this method to delete a sticker set that was created by the bot. Returns True on success.
func (b *Bot) DeleteStickerSet(ctx context.Context, params params.DeleteStickerSet) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "deleteStickerSet", params)
}

// --- Inline modes --- //

// AnswerInlineQuery Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
func (b *Bot) AnswerInlineQuery(ctx context.Context, params params.AnswerInlineQuery) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "answerInlineQuery", params)
}

// AnswerWebAppQuery Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a models.SentWebAppMessage object is returned.
func (b *Bot) AnswerWebAppQuery(ctx context.Context, params params.AnswerWebAppQuery) (models.SentWebAppMessage, error) {
	return network.UseMethod[models.SentWebAppMessage](b.client, ctx, "answerWebAppQuery", params)
}

// SendInvoice Use this method to send invoices. On success, the sent models.Message is returned.
func (b *Bot) SendInvoice(ctx context.Context, params params.SendInvoice) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendInvoice", params)
}

// CreateInvoiceLink Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (b *Bot) CreateInvoiceLink(ctx context.Context, params params.CreateInvoiceLink) (string, error) {
	return network.UseMethod[string](b.client, ctx, "createInvoiceLink", params)
}

// AnswerShippingQuery Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (b *Bot) AnswerShippingQuery(ctx context.Context, params params.AnswerShippingQuery) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "answerShippingQuery", params)
}

// AnswerPreCheckoutQuery Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an models.Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, params params.AnswerPreCheckoutQuery) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "answerPreCheckoutQuery", params)
}

// GetMyStarBalance A method to get the current Telegram Stars balance of the bot. Requires no parameters. On success, returns a models.StarAmount object.
func (b *Bot) GetMyStarBalance(ctx context.Context) (models.StarAmount, error) {
	return network.UseMethod[models.StarAmount](b.client, ctx, "getMyStarBalance", nil)
}

// GetStarTransactions Returns the bot's Telegram Star transactions in chronological order. On success, returns a models.StarTransactions object.
func (b *Bot) GetStarTransactions(ctx context.Context, params params.GetStarTransactions) (models.StarTransactions, error) {
	return network.UseMethod[models.StarTransactions](b.client, ctx, "getStarTransactions", params)
}

// RefundStarPayment Refunds a successful payment in Telegram Stars. Returns True on success.
func (b *Bot) RefundStarPayment(ctx context.Context, params params.RefundStarPayment) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "refundStarPayment", params)
}

// EditUserStarSubscription Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
func (b *Bot) EditUserStarSubscription(ctx context.Context, params params.EditUserStarSubscription) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "editUserStarSubscription", params)
}

// --- Passport --- //

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (b *Bot) SetPassportDataErrors(ctx context.Context, params params.SetPassportDataErrors) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setPassportDataErrors", params)
}

// --- Games --- //

// SendGame Use this method to send a game. On success, the sent models.Message is returned.
func (b *Bot) SendGame(ctx context.Context, params params.SendGame) (models.Message, error) {
	return network.UseMethod[models.Message](b.client, ctx, "sendGame", params)
}

// SetGameScore Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (b *Bot) SetGameScore(ctx context.Context, params params.SetGameScore) (bool, error) {
	return network.UseMethod[bool](b.client, ctx, "setGameScore", params)
}

// GetGameHighScores Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of models.GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
func (b *Bot) GetGameHighScores(ctx context.Context, params params.GetGameHighScores) ([]models.GameHighScore, error) {
	return network.UseMethod[[]models.GameHighScore](b.client, ctx, "getGameHighScores", params)
}

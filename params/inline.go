package params

import "github.com/gram-kit/gramkit-go/models"

type AnswerInlineQuery struct {
	InlineQueryID string                            `json:"inline_query_id"`
	Results       []models.InlineQueryResult        `json:"results"`
	CacheTime     int                               `json:"cache_time,omitempty"`
	IsPersonal    bool                              `json:"is_personal,omitempty"`
	NextOffset    string                            `json:"next_offset,omitempty"`
	Buttons       []models.InlineQueryResultsButton `json:"buttons,omitempty"`
}

type AnswerWebAppQuery struct {
	WebAppQueryID string                   `json:"web_app_query_id"`
	Result        models.InlineQueryResult `json:"result"`
}

type SavePreparedInlineMessage struct {
	UserID            int64                    `json:"user_id"`
	Result            models.InlineQueryResult `json:"result"`
	AllowUserChats    bool                     `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool                     `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool                     `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool                     `json:"allow_channel_chats,omitempty"`
}

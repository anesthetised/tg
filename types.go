package tg

import (
	"encoding/json"
	"time"
)

type Update struct {
	ID      int64    `json:"update_id"`
	Message *Message `json:"message"`
}

type Message struct {
	ID               int64  `json:"message_id"`
	Unixtime         int64  `json:"date"`
	Chat             Chat   `json:"chat"`
	From             *User  `json:"from"`
	Text             string `json:"text"`
	GroupChatCreated bool   `json:"group_chat_created"`
	NewChatTitle     string `json:"new_chat_title,omitempty"`
	NewChatMembers   []User `json:"new_chat_members,omitempty"`
	LeftChatMember   *User  `json:"left_chat_member,omitempty"`
}

func (m *Message) Time() time.Time {
	return time.Unix(m.Unixtime, 0)
}

func (m *Message) Sender() *User {
	if m == nil {
		return nil
	}

	switch {
	case m.From != nil:
		return m.From
	}

	return nil
}

func (m *Message) Method() string {
	return "sendMessage"
}

func (m *Message) Params() json.RawMessage {
	params, _ := json.Marshal(m)
	return params
}

type Chat struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	Type  ChatType `json:"type"`
}

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type User struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	LanguageCode  string `json:"language_code"`
	IsBot         bool   `json:"is_bot"`
	CanJoinGroups bool   `json:"can_join_groups"`
}

type RequestGetUpdates struct {
	Offset         int64    `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (RequestGetUpdates) Method() string {
	return "getUpdates"
}

func (r RequestGetUpdates) Params() json.RawMessage {
	params, _ := json.Marshal(r)
	return params
}

type requestGetMe struct{}

func (r requestGetMe) Method() string {
	return "getMe"
}

func (r requestGetMe) Params() json.RawMessage {
	return nil
}

type MessageConfig[R Recipient, M ReplyMarkup] struct {
	ChatID              int64               `json:"chat_id"`
	ThreadID            int64               `json:"message_thread_id,omitempty"`
	Text                string              `json:"text"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	Entities            []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions  *LinkPreviewOptions `json:"LinkPreviewOptions,omitempty"`
	DisableNotification bool                `json:"disable_notification"`
	ProtectContent      bool                `json:"protect_content"`
	ReplyParameters     *ReplyParameters[R] `json:"reply_parameters,omitempty"`
	ReplyMarkup         M                   `json:"reply_markup,omitempty"`
}

func (c MessageConfig[R, M]) Method() string {
	return "sendMessage"
}

func (c MessageConfig[R, M]) Params() json.RawMessage {
	params, _ := json.Marshal(c)
	return params
}

type ParseMode string

const (
	ParseModeHTML       ParseMode = "HTML"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeMarkdown   ParseMode = "Markdown"
)

type MessageEntity struct {
	Type          MessageEntityType `json:"type"`
	Offset        int64             `json:"offset"`
	Length        int64             `json:"length"`
	URL           string            `json:"url"`
	User          *User             `json:"user,omitempty"`
	Language      string            `json:"language,omitempty"`
	CustomEmojiID string            `json:"custom_emoji_id,omitempty"`
}

type MessageEntityType string

const (
	MessageEntityyMention      MessageEntityType = "mention"
	MessageEntityHashtag       MessageEntityType = "hashtag"
	MessageEntityCashtag       MessageEntityType = "cashtag"
	MessageEntityBotCommand    MessageEntityType = "bot_command"
	MessageEntityURL           MessageEntityType = "url"
	MessageEntityEmail         MessageEntityType = "email"
	MessageEntityPhoneNumber   MessageEntityType = "phone_number"
	MessageEntityBold          MessageEntityType = "bold"
	MessageEntityItalic        MessageEntityType = "italic"
	MessageEntityUnderlime     MessageEntityType = "underline"
	MessageEntityStrikethrough MessageEntityType = "strikethrough"
	MessageEntitySpoiler       MessageEntityType = "spoiler"
	MessageEntityBlockquote    MessageEntityType = "blockquote"
	MessageEntityCode          MessageEntityType = "code"
	MessageEntityPre           MessageEntityType = "pre"
	MessageEntityTextLink      MessageEntityType = "text_link"
	MessageEntityTextMention   MessageEntityType = "text_mention"
	MessageEntityCustomEmoji   MessageEntityType = "custom_emoji"
)

type LinkPreviewOptions struct {
	Disabled         bool   `json:"is_disabled"`
	URL              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media"`
	PreferLargeMedia bool   `json:"prefer_large_media"`
	ShowAboveText    bool   `json:"show_above_text"`
}

type ReplyParameters[T Recipient] struct {
	MessageID int64 `json:"message_id"`
	ChatID    T     `json:"chat_id"`
}

type Recipient interface {
	~int64 | ~string
}

type ReplyMarkup interface {
	KeyboardMarkup | InlineKeyboardMarkup
}

type KeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

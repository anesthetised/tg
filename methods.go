package tg

import (
	"context"
)

func GetUpdates(ctx context.Context, c Caller, req RequestGetUpdates) ([]Update, error) {
	resp, err := c.Call(ctx, req)
	if err != nil {
		return nil, err
	}
	return decodePayload[[]Update](resp)
}

func GetMe(ctx context.Context, c Caller) (*User, error) {
	resp, err := c.Call(ctx, RequestGetMe{})
	if err != nil {
		return nil, err
	}
	return decodePayload[*User](resp)
}

func SendMessage[R Recipient, M ReplyMarkup](ctx context.Context,
	c Caller,
	chatID int64,
	text string,
	opts ...OptionMessage[R, M],
) (*Message, error) {
	conf := MessageConfig[R, M]{
		ChatID: chatID,
		Text:   text,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	resp, err := c.Call(ctx, conf)
	if err != nil {
		return nil, err
	}

	return decodePayload[*Message](resp)
}

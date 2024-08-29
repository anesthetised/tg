package tg

import (
	"context"
	"time"
)

type OnError func(req RequestGetUpdates, err error) (stop bool)

func Updates(ctx context.Context, c Caller, limit int, timeout time.Duration, onErr OnError) <-chan Update {
	updatesChan := make(chan Update, 1)

	go func() {
		defer close(updatesChan)

		req := RequestGetUpdates{
			Limit:   limit,
			Timeout: int(timeout.Microseconds() * 1000),
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			updates, err := GetUpdates(ctx, c, req)
			if err != nil {
				if onErr != nil && onErr(req, err) {
					break
				}

				continue
			}

			for _, update := range updates {
				req.Offset = update.ID + 1
				updatesChan <- update
			}
		}
	}()

	return updatesChan
}

type OptionMessage[R Recipient, M ReplyMarkup] func(c *MessageConfig[R, M])

func NewKeyboardButtonRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

func NewKeyboard(rows ...[]KeyboardButton) KeyboardMarkup {
	markup := KeyboardMarkup{
		Keyboard:        make([][]KeyboardButton, len(rows)),
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	for i, row := range rows {
		markup.Keyboard[i] = row
	}

	return markup
}

func NewInlineKeyboardButtonRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

func NewInlineKeyboard(rows ...[]InlineKeyboardButton) InlineKeyboardMarkup {
	markup := InlineKeyboardMarkup{
		InlineKeyboard: make([][]InlineKeyboardButton, len(rows)),
	}

	for i, row := range rows {
		markup.InlineKeyboard[i] = row
	}

	return markup
}

func WithKeyboard[R Recipient](markup KeyboardMarkup) OptionMessage[R, KeyboardMarkup] {
	return func(c *MessageConfig[R, KeyboardMarkup]) {
		c.ReplyMarkup = markup
	}
}

func WithInlineKeyboard[R Recipient](markup InlineKeyboardMarkup) OptionMessage[R, InlineKeyboardMarkup] {
	return func(c *MessageConfig[R, InlineKeyboardMarkup]) {
		c.ReplyMarkup = markup
	}
}

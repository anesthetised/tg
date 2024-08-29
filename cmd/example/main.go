package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anesthetised/tg"
)

func run(ctx context.Context, conf *Config, logger *slog.Logger) error {
	var err error

	if err = conf.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	caller := tg.New(conf.Token,
		tg.WithBaseURL(conf.BaseURL),
		tg.WithLogger(logger),
	)

	user, err := tg.GetMe(ctx, caller)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	logger.Info("user info",
		"id", user.ID,
		"username", user.Username,
		"first_name", user.FirstName,
		"last_name", user.LastName,
		"is_bot", user.IsBot,
	)

	for upd := range tg.Updates(ctx, caller, 0, 0, nil) {
		msg := upd.Message
		if msg == nil {
			continue
		}

		sender := msg.Sender()

		logger.Info("message",
			"id", msg.ID,
			"time", msg.Time(),
			slog.Group("chat",
				"id", msg.Chat.ID,
				"type", msg.Chat.Type,
				"title", msg.Chat.Title,
			),
			slog.Group("sender",
				"id", sender.ID,
				"username", sender.Username,
				"first_name", sender.FirstName,
				"last_name", sender.LastName,
				"lang_code", sender.LanguageCode,
			),
			"text", msg.Text,
		)
	}

	return nil
}

func main() {
	var conf Config

	flag.StringVar(&conf.BaseURL, "url", "", "API base URL")
	flag.StringVar(&conf.Token, "token", "", "bot token")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := run(ctx, &conf, logger)
	if errors.Is(err, context.Canceled) {
		logger.Info("canceled")
		return
	}
	if err != nil {
		logger.Error("operation is failed", "err", err)
		return
	}

	logger.Info("done")
}

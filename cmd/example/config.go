package main

import "errors"

type Config struct {
	BaseURL string
	Token   string
}

func (c Config) Validate() error {
	if c.Token == "" {
		return errors.New("token is required")
	}

	return nil
}

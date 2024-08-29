package tg

import (
	"context"
	"encoding/json"
)

type clientMock[T any] struct {
	value T
}

func newClientMock[T any](value T) *clientMock[T] {
	return &clientMock[T]{
		value: value,
	}
}

func (c *clientMock[T]) Call(_ context.Context, _ Sendable) (json.RawMessage, error) {
	return json.Marshal(c.value)
}

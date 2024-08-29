package tg

import (
	"context"
	"encoding/json"
)

type Sendable interface {
	Method() string
	Params() json.RawMessage
}

type Caller interface {
	Call(ctx context.Context, value Sendable) (json.RawMessage, error)
}

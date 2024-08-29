package tg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ Caller = (*Client)(nil)
	_ Caller = (*clientMock[any])(nil)
)

func TestClient(t *testing.T) {
	zero := requestGetMe{}
	c := newClientMock[requestGetMe](zero)

	data, err := c.Call(context.Background(), zero)
	assert.NoError(t, err)

	req, err := decodePayload[requestGetMe](data)
	assert.NoError(t, err)

	assert.ObjectsAreEqual(zero, req)
}

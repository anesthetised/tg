package tg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUpdates(t *testing.T) {
	const expectedValue = 1

	c := newClientMock[[]Update]([]Update{
		{
			ID: expectedValue,
		},
	})

	updates, err := GetUpdates(context.Background(), c, RequestGetUpdates{})
	assert.NoError(t, err)
	assert.Len(t, updates, expectedValue)

	upd := updates[0]

	assert.Equal(t, int64(expectedValue), upd.ID)
	assert.Nil(t, upd.Message)
}

func TestGetMe(t *testing.T) {
	const expectedValue = "test"

	c := newClientMock[*User](&User{
		Username: expectedValue,
	})

	user, err := GetMe(context.Background(), c)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, user.Username)
}

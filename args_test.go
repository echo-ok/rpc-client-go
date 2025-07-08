package rpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs_Add(t *testing.T) {
	store := Store{
		ID:    "-1",
		Name:  "Temu SEMI Store",
		Env:   "prod",
		Debug: true,
		Configuration: Configuration{
			"region":             "",
			"app_key":            "",
			"app_secret":         "",
			"access_token":       "",
			"static_file_server": "",
		},
	}
	args := NewArgs().Add(NewPayload(store)).Add(NewPayload(store))
	assert.Equal(t, 1, len(args))
	args = NewArgs().Add(NewPayload(store)).Add(NewPayload(store).SetBody(1))
	assert.Equal(t, 2, len(args))

	store2 := store
	args = NewArgs().Add(NewPayload(store)).Add(NewPayload(store2).SetBody())
	assert.Equal(t, 1, len(args))

	args = NewArgs().Add(NewPayload(store)).Add(NewPayload(store2).SetBody(1))
	assert.Equal(t, 2, len(args))

	args = NewArgs().Add(NewPayload(store)).Add(NewPayload(store2).SetBody([]string{}))
	assert.Equal(t, 1, len(args))

	args = NewArgs().Add(NewPayload(store)).Add(NewPayload(store2).SetBody(""))
	assert.Equal(t, 1, len(args))

	args = NewArgs().
		Add(NewPayload(store)).
		Add(NewPayload(store2).SetBody("xxxx")).
		Del("-1")
	assert.Equal(t, 0, len(args))
}

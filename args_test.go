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
	assert.Equal(t, nil, args[0].Body)
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
func TestArgs_SetBody(t *testing.T) {
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
	args := NewArgs().Add(NewPayload(store))
	args = args.Add(NewPayload(store, "123"))
	assert.Equal(t, 2, len(args))
	assert.Equal(t, "123", args[1].Body)

	args = args.SetBody("abc")
	assert.Equal(t, 2, len(args))
	assert.Equal(t, "abc", args[0].Body)
	assert.Equal(t, "abc", args[1].Body)
}

func TestArgs_SetStoreBody(t *testing.T) {
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
	store1 := store
	store1.ID = "-2"
	args := NewArgs().Add(NewPayload(store))
	args = args.Add(NewPayload(store1, "123"))
	assert.Equal(t, 2, len(args))
	assert.Equal(t, nil, args[0].Body)
	assert.Equal(t, "123", args[1].Body)

	args = args.SetStoreBody("-1", "abc")
	assert.Equal(t, 2, len(args))
	assert.Equal(t, "abc", args[0].Body)
	assert.Equal(t, "123", args[1].Body)

	args = args.SetStoreBody("-2", "abc")
	assert.Equal(t, 2, len(args))
	assert.Equal(t, "abc", args[0].Body)
	assert.Equal(t, "abc", args[1].Body)
}

func TestArgs_isEmptyValue(t *testing.T) {
	// 空值
	ptrString := ""
	ptrMap := map[string]string{}
	ptrArray := []string{}
	ptrSlice := make([]any, 0)
	values := []any{
		nil,
		"",
		" ",
		"	", // \t
		[]string{},
		[]int{},
		map[string]string{},
		map[int]string{},
		map[int]int{},
		map[string]int{},
		map[string]any{},
		map[any]any{},
		map[any]string{},
		map[any]int{},
		&ptrString,
		&ptrMap,
		&ptrArray,
		&ptrSlice,
	}
	for k, v := range values {
		assert.True(t, isEmptyValue(v), "#%d", k)
	}

	// 非空值
	values = []any{
		0,
		false,
		[]string{""},
		[]any{nil},
	}
	for k, v := range values {
		assert.False(t, isEmptyValue(v), "#%d", k)
	}
}

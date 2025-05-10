package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguration_read(t *testing.T) {
	tests := Configuration{
		"a": "aa",
		"b": nil,
		"c": 1,
		"d": true,
		"e": "",
	}
	for k, v := range tests {
		assert.Equalf(t, v, tests.read(k), "%s => %s", k, v)
	}
}

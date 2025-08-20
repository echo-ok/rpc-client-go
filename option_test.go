package rpclient

import (
	"testing"
)

func TestConstants(t *testing.T) {
	if JsonCodec != "json" {
		t.Errorf("Expected JsonCodec to be 'json', got %s", JsonCodec)
	}
	if GoridgeCodec != "goridge" {
		t.Errorf("Expected GoridgeCodec to be 'goridge', got %s", GoridgeCodec)
	}
}

func TestDefaultOption(t *testing.T) {
	// Test that defaultOption has expected values
	if defaultOption.Network != "tcp" {
		t.Errorf("Expected defaultOption.Network to be 'tcp', got %s", defaultOption.Network)
	}
	if defaultOption.Codec != JsonCodec {
		t.Errorf("Expected defaultOption.Codec to be 'json', got %s", defaultOption.Codec)
	}
	if defaultOption.LogLevel != "debug" {
		t.Errorf("Expected defaultOption.LogLevel to be 'debug', got %s", defaultOption.LogLevel)
	}
	
	// Test that sensitive words contain expected values
	expectedSensitiveWords := []string{
		"key",
		"app_key",
		"china_app_key",
		"secret",
		"app_secret",
		"china_app_secret",
		"token",
		"access_token",
		"china_access_token",
		"name",
		"username",
		"pwd",
		"password",
	}
	
	if len(defaultOption.SensitiveWords) != len(expectedSensitiveWords) {
		t.Errorf("Expected %d sensitive words, got %d", len(expectedSensitiveWords), len(defaultOption.SensitiveWords))
	}
	
	for i, word := range expectedSensitiveWords {
		if i >= len(defaultOption.SensitiveWords) {
			t.Errorf("Missing sensitive word: %s", word)
			continue
		}
		if defaultOption.SensitiveWords[i] != word {
			t.Errorf("Expected sensitive word %d to be '%s', got '%s'", i, word, defaultOption.SensitiveWords[i])
		}
	}
}

func TestOptionStruct(t *testing.T) {
	// Test Option struct creation and field types
	option := Option{
		Network:        "tcp",
		Codec:          "json",
		LogLevel:       "info",
		SensitiveWords: []string{"test", "secret"},
	}
	
	if option.Network != "tcp" {
		t.Errorf("Expected Network to be 'tcp', got %s", option.Network)
	}
	if option.Codec != "json" {
		t.Errorf("Expected Codec to be 'json', got %s", option.Codec)
	}
	if option.LogLevel != "info" {
		t.Errorf("Expected LogLevel to be 'info', got %s", option.LogLevel)
	}
	
	if len(option.SensitiveWords) != 2 {
		t.Errorf("Expected 2 sensitive words, got %d", len(option.SensitiveWords))
	}
	if option.SensitiveWords[0] != "test" {
		t.Errorf("Expected first sensitive word to be 'test', got %s", option.SensitiveWords[0])
	}
	if option.SensitiveWords[1] != "secret" {
		t.Errorf("Expected second sensitive word to be 'secret', got %s", option.SensitiveWords[1])
	}
}

func TestOptionWithEmptyValues(t *testing.T) {
	// Test Option with empty values
	option := Option{
		Network:        "",
		Codec:          "",
		LogLevel:       "",
		SensitiveWords: []string{},
	}
	
	if option.Network != "" {
		t.Errorf("Expected Network to be empty, got %s", option.Network)
	}
	if option.Codec != "" {
		t.Errorf("Expected Codec to be empty, got %s", option.Codec)
	}
	if option.LogLevel != "" {
		t.Errorf("Expected LogLevel to be empty, got %s", option.LogLevel)
	}
	
	if len(option.SensitiveWords) != 0 {
		t.Errorf("Expected SensitiveWords to be empty, got %d words", len(option.SensitiveWords))
	}
}

func TestOptionWithNilSensitiveWords(t *testing.T) {
	// Test Option with nil SensitiveWords
	option := Option{
		Network:  "tcp",
		Codec:    "json",
		LogLevel: "debug",
	}
	
	// SensitiveWords should be nil by default
	if option.SensitiveWords != nil {
		t.Errorf("Expected SensitiveWords to be nil, got %v", option.SensitiveWords)
	}
}

func TestOptionCodecTypes(t *testing.T) {
	// Test Option with different codec types
	codecs := []string{JsonCodec, GoridgeCodec}
	
	for _, codec := range codecs {
		option := Option{
			Network:  "tcp",
			Codec:    codec,
			LogLevel: "debug",
		}
		if option.Codec != codec {
			t.Errorf("Expected Codec to be '%s', got '%s'", codec, option.Codec)
		}
	}
}

func TestOptionNetworkTypes(t *testing.T) {
	// Test Option with different network types
	networks := []string{"tcp", "tcp4", "tcp6", "unix"}
	
	for _, network := range networks {
		option := Option{
			Network:  network,
			Codec:    "json",
			LogLevel: "debug",
		}
		if option.Network != network {
			t.Errorf("Expected Network to be '%s', got '%s'", network, option.Network)
		}
	}
}

func TestOptionLogLevelTypes(t *testing.T) {
	// Test Option with different log levels
	logLevels := []string{"debug", "info", "warn", "error"}
	
	for _, level := range logLevels {
		option := Option{
			Network:  "tcp",
			Codec:    "json",
			LogLevel: level,
		}
		if option.LogLevel != level {
			t.Errorf("Expected LogLevel to be '%s', got '%s'", level, option.LogLevel)
		}
	}
}

func TestOptionSensitiveWordsModification(t *testing.T) {
	// Test that SensitiveWords can be modified
	option := Option{
		Network:        "tcp",
		Codec:          "json",
		LogLevel:       "debug",
		SensitiveWords: []string{"initial"},
	}
	
	// Add new sensitive word
	option.SensitiveWords = append(option.SensitiveWords, "new_secret")
	found := false
	for _, word := range option.SensitiveWords {
		if word == "new_secret" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected 'new_secret' to be in SensitiveWords")
	}
	
	// Replace entire list
	option.SensitiveWords = []string{"completely_new", "another_secret"}
	if len(option.SensitiveWords) != 2 {
		t.Errorf("Expected 2 sensitive words, got %d", len(option.SensitiveWords))
	}
	if option.SensitiveWords[0] != "completely_new" {
		t.Errorf("Expected first sensitive word to be 'completely_new', got %s", option.SensitiveWords[0])
	}
	if option.SensitiveWords[1] != "another_secret" {
		t.Errorf("Expected second sensitive word to be 'another_secret', got %s", option.SensitiveWords[1])
	}
}
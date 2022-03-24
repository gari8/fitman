package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIoHandler(t *testing.T) {
	assert.Equal(t, &IoHandler{}, NewIoHandler())
}

func TestIoHandler_DecodeToml(t *testing.T) {
	expected := TomlSetting{
		Config{
			ApiKey: "test",
		},
	}
	var actual TomlSetting
	to := `[default]
api_key="test"`
	io := NewIoHandler()
	_, err := io.DecodeToml(to, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

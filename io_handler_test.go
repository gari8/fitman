package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIoHandler(t *testing.T) {
	assert.Equal(t, &IoHandler{}, NewIoHandler())
}

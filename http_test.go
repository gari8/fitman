package main

import (
	"bytes"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	assert.Equal(t, &HttpClient{}, NewHttpClient())
}

func TestHttpClient_post(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testPath := "http://localhost:1111"

	httpmock.RegisterResponder(
		"POST",
		"http://localhost:1111",
		httpmock.NewStringResponder(200, "OK"))

	c := NewHttpClient()
	b, err := c.post(testPath, nil, nil)
	assert.NoError(t, err)
	expected, err := ioutil.ReadAll(bytes.NewBuffer([]byte("OK")))
	assert.NoError(t, err)
	assert.Equal(t, expected, b)
}

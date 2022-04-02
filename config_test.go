package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getMockIoHandler(t *testing.T) *MockioHandler {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockioHandler(ctrl)
}

func getMockHttpClient(t *testing.T) *MockhttpClient {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockhttpClient(ctrl)
}

func TestNewConfig(t *testing.T) {
	io := getMockIoHandler(t)
	hc := getMockHttpClient(t)
	assert.Equal(t, &Config{
		ioHandler:  io,
		httpClient: hc,
	}, NewConfig(hc, io))
}

func TestConfig_readConfig(t *testing.T) {
	toml := `[default]
name=test`
	io := getMockIoHandler(t)
	io.EXPECT().DecodeToml(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	io.EXPECT().GetHomeDirPath().Return("/", nil).AnyTimes()
	hc := getMockHttpClient(t)
	c := NewConfig(hc, io)
	c, err := c.readConfig([]byte(toml))
	assert.NoError(t, err)
}

func TestConfig_refresh(t *testing.T) {
	expected := []byte("test")
	io := getMockIoHandler(t)
	io.EXPECT().MakeDir(gomock.Any()).Return(nil).AnyTimes()
	io.EXPECT().OpenFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(&os.File{}, nil).AnyTimes()
	io.EXPECT().Write(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
	io.EXPECT().ReadFile(gomock.Any()).Return(expected, nil).AnyTimes()
	io.EXPECT().DecodeToml(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	io.EXPECT().NotExists(gomock.Any()).Return(false).AnyTimes()
	io.EXPECT().GetHomeDirPath().Return("/", nil).AnyTimes()
	hc := getMockHttpClient(t)
	hc.EXPECT().post(gomock.Any(), gomock.Any(), gomock.Any()).Return(expected, nil).AnyTimes()
	c := NewConfig(hc, io)

	b, err := c.refresh()
	assert.NoError(t, err)
	assert.Equal(t, expected, b)
}

func TestConfig_setConfigFile(t *testing.T) {
	expected := []byte(`{
"kind": "a",
"idToken": "a",
"refreshToken": "a",
"expiresIn": "a",
"localId": "a"
}`)
	io := getMockIoHandler(t)
	io.EXPECT().NotExists(gomock.Any()).Return(false).AnyTimes()
	io.EXPECT().GetHomeDirPath().Return("/", nil).AnyTimes()
	io.EXPECT().MakeDir(gomock.Any()).Return(nil).AnyTimes()
	io.EXPECT().OpenFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(&os.File{}, nil).AnyTimes()
	io.EXPECT().Write(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
	hc := getMockHttpClient(t)
	hc.EXPECT().post(gomock.Any(), gomock.Any(), gomock.Any()).Return(expected, nil).AnyTimes()
	c := NewConfig(hc, io)
	c.ApiKey = "test"

	b, err := c.setConfigFile()
	assert.NoError(t, err)
	assert.Equal(t, expected, b)
}

func TestConfig_writeFiles(t *testing.T) {
	io := getMockIoHandler(t)
	io.EXPECT().NotExists(gomock.Any()).Return(false).AnyTimes()
	io.EXPECT().GetHomeDirPath().Return("/", nil).AnyTimes()
	io.EXPECT().MakeDir(gomock.Any()).Return(nil).AnyTimes()
	io.EXPECT().OpenFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(&os.File{}, nil).AnyTimes()
	io.EXPECT().Write(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
	hc := getMockHttpClient(t)
	c := NewConfig(hc, io)
	c.ApiKey = "test"

	assert.NoError(t, c.writeFiles())
}

func TestConfig_find(t *testing.T) {
	io := getMockIoHandler(t)
	io.EXPECT().NotExists(gomock.Any()).Return(false).AnyTimes()
	io.EXPECT().GetHomeDirPath().Return("/", nil).AnyTimes()
	io.EXPECT().ReadFile(gomock.Any()).Return([]byte("test"), nil).AnyTimes()
	hc := getMockHttpClient(t)
	c := NewConfig(hc, io)
	c.ApiKey = "test"
	body, err := c.find()
	assert.NoError(t, err)
	assert.Equal(t, body, []byte("test"))
}

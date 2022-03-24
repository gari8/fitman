package main

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./mock.go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	basedir                 = ".fitman"
	configFile              = "config.toml"
	firebaseEndpoint        = "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s"
	firebaseRefreshEndpoint = "https://securetoken.googleapis.com/v1/token?key=%s"
	fileContentFormat       = `[default]
api_key="%s"
refresh_token="%s"`
)

type tokenInfo struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

type TomlSetting struct {
	Config `toml:"default"`
}

type httpClient interface {
	post(path string, value io.Reader, header map[string]string) ([]byte, error)
}

type ioHandler interface {
	ReadFile(path string) ([]byte, error)
	DecodeToml(data string, v *TomlSetting) (interface{}, error)
	MakeDir(dirPath string) error
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Write(f *os.File, b []byte) (n int, err error)
}

type Config struct {
	ApiKey       string `toml:"api_key"`
	RefreshToken string `toml:"refresh_token"`
	httpClient
	ioHandler
}

func NewConfig(c httpClient,
	h ioHandler) *Config {
	return &Config{
		httpClient: c,
		ioHandler:  h,
	}
}

func (c Config) readConfig() (*Config, error) {
	fullPath := basedir + "/" + configFile
	content, err := c.ioHandler.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	var t TomlSetting
	_, err = c.ioHandler.DecodeToml(string(content), &t)
	if err != nil {
		return nil, err
	}
	t.Config.httpClient = c.httpClient
	t.Config.ioHandler = c.ioHandler
	return &t.Config, nil
}

func (c *Config) refresh() ([]byte, error) {
	c, err := c.readConfig()
	if err != nil {
		return nil, err
	}
	body, err := c.httpClient.post(fmt.Sprintf(firebaseRefreshEndpoint, c.ApiKey),
		bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", c.RefreshToken))),
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Config) setConfigFile() ([]byte, error) {
	fullPath := basedir + "/" + configFile

	if err := c.ioHandler.MakeDir(basedir); err != nil {
		return nil, err
	}

	f, err := c.ioHandler.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	if c.ApiKey == "" {
		return nil, errors.New("APIKEY is not found")
	}

	body, err := c.httpClient.post(fmt.Sprintf(firebaseEndpoint, c.ApiKey), nil, map[string]string{
		"Content-Type": "application/json",
	})

	if err != nil {
		return nil, err
	}
	var info tokenInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, err
	}

	c.RefreshToken = info.RefreshToken

	if _, err := c.ioHandler.Write(f, []byte(fmt.Sprintf(fileContentFormat,
		c.ApiKey, c.RefreshToken))); err != nil {
		return nil, err
	}

	return body, nil
}

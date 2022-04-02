package main

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./mock.go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	basedir                 = ".fitman"
	configFile              = "config.toml"
	firebaseEndpoint        = "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s"
	firebaseRefreshEndpoint = "https://securetoken.googleapis.com/v1/token?key=%s"
	fileContentFormat       = `
[%s]
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

type httpClient interface {
	post(path string, value io.Reader, header map[string]string) ([]byte, error)
}

type ioHandler interface {
	ReadFile(path string) ([]byte, error)
	DecodeToml(data string, v interface{}) (interface{}, error)
	MakeDir(dirPath string) error
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Write(f *os.File, b []byte) (n int, err error)
	NotExists(path string) bool
	GetHomeDirPath() (string, error)
	RemoveFile(path string) error
}

type Config struct {
	ApiKey       string `toml:"api_key"`
	RefreshToken string `toml:"refresh_token"`
	Profile      string `json:"profile"`
	SubCmd       string `json:"subCmd"`
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

func (c Config) readConfig(body []byte) (*Config, error) {
	var t map[string]interface{}
	_, err := c.ioHandler.DecodeToml(string(body), &t)
	if err != nil {
		return nil, err
	}
	m, ok := t[c.Profile].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no field %s", c.Profile)
	}
	c.ApiKey = m["api_key"].(string)
	c.RefreshToken = m["refresh_token"].(string)
	return &c, nil
}

func (c Config) contains(body []byte, key string) (bool, error) {
	var t map[string]interface{}
	_, err := c.ioHandler.DecodeToml(string(body), &t)
	if err != nil {
		return false, err
	}
	for k, _ := range t {
		if k == key {
			return true, nil
		}
	}
	return false, nil
}

func (c *Config) refresh() ([]byte, error) {
	body, err := c.httpClient.post(fmt.Sprintf(firebaseRefreshEndpoint, c.ApiKey),
		bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", c.RefreshToken))),
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Config) setConfigFile() ([]byte, error) {
	if c.ApiKey == "" {
		return nil, errors.New("apiKey is not found")
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
	return body, c.writeFiles()
}

func (c Config) writeFiles() error {
	home, err := c.ioHandler.GetHomeDirPath()
	if err != nil {
		return err
	}
	if c.ioHandler.NotExists(filepath.Join(home, basedir)) {
		if err := c.ioHandler.MakeDir(filepath.Join(home, basedir)); err != nil {
			return err
		}
	}
	if c.SubCmd == initialization {
		_ = c.ioHandler.RemoveFile(filepath.Join(home, basedir, configFile))
	}
	f, err := c.ioHandler.OpenFile(filepath.Join(home, basedir, configFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	if _, err := c.ioHandler.Write(f, []byte(fmt.Sprintf(fileContentFormat,
		c.Profile, c.ApiKey, c.RefreshToken))); err != nil {
		return err
	}
	return nil
}

func (c Config) find() ([]byte, error) {
	home, err := c.ioHandler.GetHomeDirPath()
	if err != nil {
		return nil, err
	}
	if c.ioHandler.NotExists(filepath.Join(home, basedir)) {
		return nil, errors.New(".fitman is not found. please `fitman init`")
	}
	if c.ioHandler.NotExists(filepath.Join(home, basedir, configFile)) {
		return nil, errors.New("setting file is not found. please `fitman init`")
	}
	return c.ioHandler.ReadFile(filepath.Join(home, basedir, configFile))
}

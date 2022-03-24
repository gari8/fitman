package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
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
	IdToken      string `json:"idToke"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}

type TomlSetting struct {
	Config `toml:"default"`
}

type Config struct {
	ApiKey       string `toml:"api_key"`
	RefreshToken string `toml:"refresh_token"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) readConfig() (*Config, error) {
	fullPath := basedir + "/" + configFile
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	var t TomlSetting
	_, err = toml.Decode(string(content), &t)
	if err != nil {
		return nil, err
	}
	return &t.Config, nil
}

func (c *Config) refresh() ([]byte, error) {
	c, err := c.readConfig()
	if err != nil {
		return nil, err
	}
	body, err := post(fmt.Sprintf(firebaseRefreshEndpoint, c.ApiKey),
		bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", c.RefreshToken))),
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Config) setConfigFile() ([]byte, error) {
	fullPath := basedir + "/" + configFile

	if _, err := os.Stat(basedir); os.IsNotExist(err) {
		if err := os.Mkdir(basedir, 0777); err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	if c.ApiKey == "" {
		return nil, errors.New("APIKEY is not found")
	}

	body, err := post(fmt.Sprintf(firebaseEndpoint, c.ApiKey), nil, map[string]string{
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

	if _, err := f.Write([]byte(fmt.Sprintf(fileContentFormat,
		c.ApiKey, c.RefreshToken))); err != nil {
		return nil, err
	}

	return body, nil
}

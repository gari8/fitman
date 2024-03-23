package modules

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/BurntSushi/toml"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

const (
	basedir           = ".fitman"
	configFile        = "config.toml"
	fileContentFormat = `
[%s]
api_key="%s"
refresh_token="%s"`
	osPerm = 0777
)

type (
	FS             struct{}
	ConfigResponse map[string]struct {
		ApiKey       string `toml:"api_key" json:"apiKey"`
		RefreshToken string `toml:"refresh_token" json:"refreshToken"`
	}
	DialogueResponse struct {
		ApiKey       string `json:"apiKey"`
		RefreshToken string `json:"refreshToken"`
	}
	ConfigRequest struct {
		Profile      string `json:"profile"`
		ApiKey       string `json:"apiKey"`
		RefreshToken string `json:"refreshToken"`
	}
)

func NewFS() *FS {
	return &FS{}
}

func (s *FS) ReadConfig() (ConfigResponse, error) {
	var resp ConfigResponse
	_, err := toml.DecodeFile(filepath.Join(homeDir, basedir, configFile), &resp)
	if err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			return nil, fmt.Errorf("")
		}
		return nil, err
	}
	return resp, nil
}

func (s *FS) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (s *FS) SetConfig(requests []ConfigRequest, override bool) error {
	basePath := filepath.Join(homeDir, basedir)
	if !s.Exists(basePath) {
		if err := os.Mkdir(basePath, osPerm); err != nil {
			return err
		}
	}
	configPath := filepath.Join(homeDir, basedir, configFile)
	flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	if !s.Exists(configPath) || override {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	f, err := os.OpenFile(configPath, flag, osPerm)
	defer f.Close()
	for _, req := range requests {
		if _, err := f.Write([]byte(fmt.Sprintf(fileContentFormat, req.Profile, req.ApiKey, req.RefreshToken))); err != nil {
			return err
		}
	}
	return err
}

func (s *FS) Dialogue(override bool) (DialogueResponse, error) {
	var resp DialogueResponse
	var answers = struct {
		ApiKey       string `survey:"apiKey"`
		RefreshToken string `survey:"refreshToken"`
		Confirm      bool   `survey:"confirm"`
	}{}
	question := qs
	if override {
		question = append(question, &survey.Question{
			Name: "confirm",
			Prompt: &survey.Confirm{
				Message: "Are you sure you want to reset all profile information?",
			},
		})
	}
	if err := survey.Ask(question, &answers); err != nil {
		return resp, err
	}
	if override && !answers.Confirm {
		return resp, fmt.Errorf("...bye")
	}
	resp.ApiKey = answers.ApiKey
	resp.RefreshToken = answers.RefreshToken
	return resp, nil
}

func (s *FS) Confirm() error {
	var answers = struct {
		Confirm bool `survey:"confirm"`
	}{}
	if err := survey.Ask([]*survey.Question{
		{
			Name: "confirm",
			Prompt: &survey.Confirm{
				Message: "Are you sure you want to reset profile information?",
			},
		},
	}, &answers); err != nil {
		return err
	}
	if !answers.Confirm {
		return fmt.Errorf("...bye")
	}
	return nil
}

func (cr ConfigResponse) Keys() []string {
	var resp []string
	for k := range cr {
		resp = append(resp, k)
	}
	slices.Sort[[]string](resp)
	return resp
}

func (cr ConfigResponse) Contains(key string) bool {
	for k := range cr {
		if k == key {
			return true
		}
	}
	return false
}

var qs = []*survey.Question{
	{
		Name: "apiKey",
		Prompt: &survey.Input{
			Message: "Please your API_KEY [required]",
			Help:    "see 'https://console.firebase.google.com/u/0/project/<your-project-name>/settings/general'",
		},
		Validate: survey.Required,
	},
	{
		Name:   "refreshToken",
		Prompt: &survey.Input{Message: "Please your REFRESH_TOKEN [optional]"},
	},
}

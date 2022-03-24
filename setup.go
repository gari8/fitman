package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

const (
	initialization = "init"
	getTokenInfo   = "get"
	version        = "version"
	help           = "help"
	helpContent    = `
[sub commands]
// create .fitman directory & get idToken
fitman init


// show idToken (after init) 
fitman get

// show help
fitman help

// show version
fitman version

[option]
fitman -v get`
)

func (c *Config) setup() {
	var verbose bool
	var m map[string]interface{}
	flag.BoolVar(&verbose, "v", false, "view detailed information")
	flag.Parse()
	subCmd := flag.Arg(0)
	switch subCmd {
	case initialization:
		var b []byte
		var err error
		if err := c.dialogue(); err != nil {
			log.Fatalln(err)
		}
		if c.RefreshToken == "" {
			b, err = c.setConfigFile()
		} else {
			b, err = c.refresh()
		}
		if err != nil {
			log.Fatalln(err)
		}
		if verbose {
			fmt.Println(string(b))
		} else {
			err := json.Unmarshal(b, &m)
			if err != nil {
				log.Fatalln(err)
			}
			// こっちはcamel case
			fmt.Println(m["idToken"])
		}
	case getTokenInfo:
		b, err := c.refresh()
		if err != nil {
			log.Fatalln(err)
		}
		if verbose {
			fmt.Println(string(b))
		} else {
			err := json.Unmarshal(b, &m)
			if err != nil {
				log.Fatalln(err)
			}
			// こっちはsnake case
			fmt.Println(m["id_token"])
		}
	case version:
		fmt.Println("v0.0.1")
	case help:
		fmt.Println(helpContent)
	}
}

var qs = []*survey.Question{
	{
		Name: "apiKey",
		Prompt: &survey.Input{
			Message: "Please your API_KEY [required]",
			Help:    "see 'https://console.firebase.google.com/u/0/project/<your-project-name>/settings/general'",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "refreshToken",
		Prompt:    &survey.Input{Message: "Please your REFRESH_TOKEN"},
		Transform: survey.Title,
	},
	{
		Name: "confirm",
		Prompt: &survey.Confirm{
			Message: "May I try to communicate with identityAPI?",
		},
	},
}

func (c *Config) dialogue() error {
	var answers = struct {
		ApiKey       string `survey:"apiKey"`
		RefreshToken string `survey:"refreshToken"`
		Confirm      bool   `survey:"confirm"`
	}{}
	if err := survey.Ask(qs, &answers); err != nil {
		return err
	}
	if !answers.Confirm {
		return errors.New("...bye")
	}
	c.ApiKey = answers.ApiKey
	c.RefreshToken = answers.RefreshToken
	return nil
}

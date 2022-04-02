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
	additional     = "add"
	getTokenInfo   = "get"
	version        = "version"
	help           = "help"
	currentVersion = "v0.4.22"
	helpContent    = `
[sub commands]
// create .fitman directory & get idToken
fitman init

// add new field "dev"
fitman add dev
// or
fitman -p dev add

// show idToken (after init) 
fitman get

// show help
fitman help

// show version
fitman version

[option]
v: verbose
	fitman -v get
	{
	  "access_token": "dummy",
	  "expires_in": "3600",
	  "token_type": "Bearer",
	  "refresh_token": "dummy",
	  "id_token": "dummy",
	  "user_id": "dummy",
	  "project_id": "dummy"
	}

p: profile
	fitman -p qa init
	fitman -p dev init
	...
	fitman -p qa get
	fitman -p dev get`
	defaultComment = `
please enter fitman help`
)

func (c *Config) setup() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "view detailed information")
	flag.StringVar(&c.Profile, "p", "default", "token profile")
	flag.Parse()
	c.SubCmd = flag.Arg(0)
	switch c.SubCmd {
	case initialization, additional:
		if c.SubCmd == additional {
			if flag.Arg(1) != "" {
				c.Profile = flag.Arg(1)
			}
			tomlBody, err := c.find()
			if err != nil {
				log.Fatalln(err)
			}
			exists, err := c.contains(tomlBody, c.Profile)
			if err != nil {
				log.Fatalln(err)
			}
			if exists {
				log.Fatalln(fmt.Errorf("`%s` key is already exist", c.Profile))
			}
		}
		var b []byte
		var err error
		if err := c.dialogue(); err != nil {
			log.Fatalln(err)
		}
		if c.RefreshToken == "" {
			b, err = c.setConfigFile()
			if err != nil {
				log.Fatalln(err)
			}
			// こっちはcamel case
			if err := c.println(verbose, b, "idToken"); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := c.writeFiles(); err != nil {
				log.Fatalln(err)
			}
			tomlBody, err := c.find()
			if err != nil {
				log.Fatalln(err)
			}
			c, err = c.readConfig(tomlBody)
			if err != nil {
				log.Fatalln(err)
			}
			b, err = c.refresh()
			if err != nil {
				log.Fatalln(err)
			}
			// こっちはsnake case
			if err := c.println(verbose, b, "id_token"); err != nil {
				log.Fatalln(err)
			}
		}
	case getTokenInfo:
		tomlBody, err := c.find()
		if err != nil {
			log.Fatalln(err)
		}
		c, err = c.readConfig(tomlBody)
		if err != nil {
			log.Fatalln(err)
		}
		b, err := c.refresh()
		if err != nil {
			log.Fatalln(err)
		}
		// こっちはsnake case
		if err := c.println(verbose, b, "id_token"); err != nil {
			log.Fatalln(err)
		}
	case version:
		fmt.Println(currentVersion)
	case help:
		fmt.Println(helpContent)
	default:
		fmt.Println(defaultComment)
	}
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
		Prompt: &survey.Input{Message: "Please your REFRESH_TOKEN"},
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

func (c Config) println(verbose bool, b []byte, key string) error {
	var m map[string]interface{}
	if verbose {
		fmt.Println(string(b))
	} else {
		err := json.Unmarshal(b, &m)
		if err != nil {
			return err
		}
		fmt.Println(m[key])
	}
	return nil
}

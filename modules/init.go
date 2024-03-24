package modules

import (
	"log"
	"os"
)

var (
	homeDir = ""
)

func init() {
	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	homeDir = path
}

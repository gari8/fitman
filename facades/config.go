package facades

import (
	"encoding/json"
	"fmt"
	"github.com/gari8/fitman/modules"
	"github.com/spf13/cobra"
)

func RunConfig(cmd *cobra.Command, args []string) error {
	fsClient := modules.NewFS()
	conf, err := fsClient.ReadConfig()
	if err != nil {
		return err
	}
	b, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

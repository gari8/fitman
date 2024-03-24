/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gari8/fitman/facades"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	cliVersion = "v1.0.0"
	onlyToken  = "only-token"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fitman",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Version: cliVersion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fitman.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.PersistentFlags().BoolP(onlyToken, "o", false, "Get only idToken")
	initCmd.PersistentFlags().BoolP(onlyToken, "o", false, "Get only idToken")
	addCmd.PersistentFlags().BoolP(onlyToken, "o", false, "Get only idToken")
}

func Handle(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var params *facades.Params
		if cmd.HasAvailablePersistentFlags() {
			onlyIdToken, err := cmd.PersistentFlags().GetBool(onlyToken)
			if err != nil {
				log.Fatalln(err)
			}
			params = &facades.Params{OnlyIdToken: onlyIdToken}
		}
		cmd.SetContext(facades.SetParams(cmd.Context(), params))
		if err := f(cmd, args); err != nil {
			log.Fatalln(err)
		}
	}
}

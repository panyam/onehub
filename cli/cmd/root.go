/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var username = ""
var password = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oh",
	Short: "The OneHub CLI",
	Long:  `The CLI for interacting with OneHub in a simpler but more flexible way`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if username == "" || password == "" {
			return errors.New(fmt.Sprintf("Invalid username '%s' or password '%s'", username, password))
		}
		return nil
	},
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
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "Username to use for basic auth for all commands")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "Password to use for basic auth for all commands")
}

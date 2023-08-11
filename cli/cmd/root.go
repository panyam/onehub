package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CmdContext struct {
	Username string
	Password string
}

// rootCmd represents the base command when called without any subcommands
func rootCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "oh",
		Short: "The OneHub CLI",
		Long:  `The CLI for interacting with OneHub in a simpler but more flexible way`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmdContext.Username == "" || cmdContext.Password == "" {
				return errors.New(fmt.Sprintf("Invalid username '%s' or password '%s'", cmdContext.Username, cmdContext.Password))
			}
			return nil
		},
	}
	out.PersistentFlags().StringVar(&cmdContext.Username, "username", "", "Username to use for basic auth for all commands")
	out.PersistentFlags().StringVar(&cmdContext.Password, "password", "", "Password to use for basic auth for all commands")
	return out
}

var rootCmd = rootCommand()
var cmdContext CmdContext

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

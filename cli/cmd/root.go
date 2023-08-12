package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

const DEFAULT_ONEHUB_HOST = "http://localhost:8080"

type CmdContext struct {
	Client *OHClient
}

var CTX = CmdContext{Client: NewOHClient("")}

// rootCmd represents the base command when called without any subcommands
func rootCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "oh",
		Short: "The OneHub CLI",
		Long:  `The CLI for interacting with OneHub in a simpler but more flexible way`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if CTX.Client.Username == "" {
				CTX.Client.Username = os.Getenv("OneHubUsername")
				if CTX.Client.Username == "" {
					return errors.New("Username not found.  Set the --username flag or the OneHubUsername environment variable")
				}
			}

			if CTX.Client.Password == "" {
				CTX.Client.Password = os.Getenv("OneHubPassword")
				if CTX.Client.Password == "" {
					return errors.New("Password not found.  Set the --password flag or the OneHubPassword environment variable")
				}
			}

			if CTX.Client.Host == "" {
				CTX.Client.Host = os.Getenv("OneHubHost")
				if CTX.Client.Host == "" {
					CTX.Client.Host = DEFAULT_ONEHUB_HOST
				}
			}
			return nil
		},
	}
	out.PersistentFlags().StringVar(&CTX.Client.Host, "host", DEFAULT_ONEHUB_HOST, "Host name to call the client against.  Envvar: OneHubHost")
	out.PersistentFlags().StringVar(&CTX.Client.Username, "username", "", "Username to use for basic auth for all commands.  Envvar: OneHubUsername")
	out.PersistentFlags().StringVar(&CTX.Client.Password, "password", "", "Password to use for basic auth for all commands.  Envvar: OneHubPassword")
	return out
}

var rootCmd = rootCommand()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

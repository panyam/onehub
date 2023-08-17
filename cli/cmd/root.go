package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

const DEFAULT_ONEHUB_HOST = "http://localhost:8080"

var Client = NewOHClient("")

// rootCmd represents the base command when called without any subcommands
func rootCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "oh",
		Short: "The OneHub CLI",
		Long:  `The CLI for interacting with OneHub in a simpler but more flexible way`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if Client.Username == "" {
				Client.Username = os.Getenv("OneHubUsername")
				if Client.Username == "" {
					return errors.New("Username not found.  Set the --username flag or the OneHubUsername environment variable")
				}
			}

			if Client.Password == "" {
				Client.Password = os.Getenv("OneHubPassword")
				if Client.Password == "" {
					return errors.New("Password not found.  Set the --password flag or the OneHubPassword environment variable")
				}
			}

			if Client.Host == "" {
				Client.Host = os.Getenv("OneHubHost")
				if Client.Host == "" {
					Client.Host = DEFAULT_ONEHUB_HOST
				}
			}
			return nil
		},
	}
	out.PersistentFlags().BoolVar(&Client.LogRequests, "log-requests", false, "Whether to log requests sent to the api server or not")
	out.PersistentFlags().StringVar(&Client.Host, "host", "", "Host name to call the client against.  Envvar: OneHubHost")
	out.PersistentFlags().StringVar(&Client.Username, "username", "", "Username to use for basic auth for all commands.  Envvar: OneHubUsername")
	out.PersistentFlags().StringVar(&Client.Password, "password", "", "Password to use for basic auth for all commands.  Envvar: OneHubPassword")
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

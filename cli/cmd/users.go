package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func usersCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "users",
		Short: "Manage users",
		Long:  `Commands to add/remove/list users in OneHub`,
	}

	out.AddCommand(getUsersCommand())
	out.AddCommand(listUsersCommand())
	out.AddCommand(deleteUsersCommand())
	out.AddCommand(newUserCommand())
	// out.AddCommand(updateUserCommand())
	return out
}

func newUserCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "new",
		Short: "Create a new user",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				return errors.New("User's name must be specified")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			avatar, _ := cmd.Flags().GetString("avatar")
			payload := StringMap{
				"user": StringMap{
					"id":     id,
					"name":   name,
					"avatar": avatar,
				},
			}
			path := "/v1/users"
			Client.Call("POST", path, nil, nil, payload)
		},
	}
	out.Flags().StringP("id", "i", "", "A custom ID to use instead of auto generating one")
	out.Flags().StringP("type", "t", "text", "Content type to assign to the content")
	out.Flags().StringP("file", "f", "", "Load user content from the given file if user not passed as a command line arg")
	out.Flags().StringP("data", "d", "", "Extra JSON data to save as part of the content")
	return out
}

func listUsersCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "list TOPIC_ID",
		Short: "List users",
		Long:  `List users in a topic`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("A TOPIC_ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			path := fmt.Sprintf("/v1/topics/%s/users", args[0])
			Client.Call("GET", path, nil, nil, nil)
		},
	}
	return out
}

func getUsersCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "get MSG_ID [...MSG_IDS]",
		Short: "Get a user for one or more user IDs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("At least one user ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, mid := range args {
				path := fmt.Sprintf("/v1/users/%s", mid)
				Client.Call("GET", path, nil, nil, nil)
			}
		},
	}
	return out
}

func deleteUsersCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "delete MSG_ID [...MSG_IDS]",
		Short: "Delete a user by one or more user IDs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("At least one user ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, mid := range args {
				path := fmt.Sprintf("/v1/users/%s", mid)
				Client.Call("DELETE", path, nil, nil, nil)
			}
		},
	}
	return out
}

func init() {
	rootCmd.AddCommand(usersCommand())
}

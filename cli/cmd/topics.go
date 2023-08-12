package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// topicsCmd represents the entire topics command hierarchy
func topicsCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "topics",
		Short: "Manage topics",
		Long:  `Group of commands to manage and interact with topic`,
	}

	out.AddCommand(getCommand())
	out.AddCommand(listCommand())
	out.AddCommand(deleteCommand())
	out.AddCommand(createCommand())
	out.AddCommand(updateCommand())
	return out
}

func listCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "list",
		Short: "List topics",
		Long:  `List topics in the system optionally filtered by name`,
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name != "" {
				CTX.Client.Call("GET", fmt.Sprintf("/v1/topics?name=%s", name), nil, nil, nil)
			} else {
				CTX.Client.Call("GET", "/v1/topics", nil, nil, nil)
			}
		},
	}
	out.Flags().StringP("name", "n", "", "Match topics with name")
	return out
}

func getCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get topics",
		Long:  `Get one or more topics by ID (or list of IDs`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Atleast one topic ID must be specified")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, topicid := range args {
				CTX.Client.Call("GET", fmt.Sprintf("/v1/topics/%s", topicid), nil, nil, nil)
			}
		},
	}
}

func deleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete topics",
		Long:  `Delete one or more topics by ID (or list of IDs`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, topicid := range args {
				CTX.Client.Call("DELETE", fmt.Sprintf("/v1/topics/%s", topicid), nil, nil, nil)
			}
		},
	}
}

func createCommand() *cobra.Command {
	out := &cobra.Command{
		Use:        "new topic_name",
		ValidArgs:  []string{"TOPIC_NAME"},
		Args:       cobra.MinimumNArgs(1),
		ArgAliases: []string{"TOPIC_NAME"},
		Short:      "Create a new topic",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "" {
				return errors.New("Invalid Topic")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetString("id")
			name := args[0]
			params := StringMap{
				"id":   id,
				"name": name,
			}
			CTX.Client.Call("POST", "/v1/topics", nil, nil, StringMap{"topic": params})
		},
	}
	out.Flags().StringP("id", "i", "", "A custom ID to use instead of auto generating one")
	return out
}

func updateCommand() *cobra.Command {
	out := &cobra.Command{
		Use:        "update",
		ArgAliases: []string{"TOPIC_ID"},
		Short:      "Update a topic",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating a topic: ", args)
		},
	}
	out.Flags().StringP("name", "n", "", "New name to set for the topic")
	return out
}

func init() {
	rootCmd.AddCommand(topicsCommand())
}

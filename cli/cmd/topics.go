package cmd

import (
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
				fmt.Println("Listing topics by name: ", name)
			} else {
				fmt.Println("Listing all topics")
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
		Run: func(cmd *cobra.Command, args []string) {
			for _, topicid := range args {
				fmt.Println("Getting topic by ID: ", topicid)
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
				fmt.Println("Deleting topic by ID: ", topicid)
			}
		},
	}
}

func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Create a new",
		Long:  `Delete one or more topics by ID (or list of IDs`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, topicid := range args {
				fmt.Println("Deleting topic by ID: ", topicid)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(topicsCommand())
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// topicsCmd represents the topics command
var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "Manage topics",
	Long:  `Group of commands to manage and interact with topic`,
}

var listCmd = &cobra.Command{
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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get topics",
	Long:  `Get one or more topics by ID (or list of IDs`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, topicid := range args {
			fmt.Println("Getting topic by ID: ", topicid)
		}
	},
}

func init() {
	rootCmd.AddCommand(topicsCmd)

	// Add the get command to our topics group
	topicsCmd.AddCommand(getCmd)

	// Add the list command to our topics group
	topicsCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("name", "n", "", "Match topics with name")
}

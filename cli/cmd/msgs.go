package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func messagesCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "msg",
		Short: "Manage messages",
		Long:  `Commands to send/receive/list messages in a topic or for a user`,
	}

	out.AddCommand(getMessagesCommand())
	out.AddCommand(listMessagesCommand())
	out.AddCommand(deleteMessagesCommand())
	out.AddCommand(sendMessageCommand())
	// out.AddCommand(updateMessageCommand())
	return out
}

func sendMessageCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "send TOPIC_ID MSG_TEXT",
		Short: "Send a message on a topic given the topic ID",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("A topic ID must be provided to send a message on")
			}
			if len(args) < 2 {
				file, _ := cmd.Flags().GetString("file")
				if file == "" {
					return errors.New("'file' flag must be passed if message text not passed via command line")
				}
				if _, err := os.ReadFile(file); err != nil {
					return err
				}
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetString("id")
			file, _ := cmd.Flags().GetString("file")
			content_type, _ := cmd.Flags().GetString("type")
			topicid := args[0]
			msg := ""
			if len(args) > 1 {
				msg = args[1]
			} else {
				contents, _ := os.ReadFile(file)
				msg = string(contents)
			}
			payload := StringMap{
				"message": StringMap{
					"id":           id,
					"content_type": content_type,
					"content_text": msg,
				},
			}
			path := fmt.Sprintf("/v1/topics/%s/messages", topicid)
			Client.Call("POST", path, nil, nil, payload)
		},
	}
	out.Flags().StringP("id", "i", "", "A custom ID to use instead of auto generating one")
	out.Flags().StringP("type", "t", "text", "Content type to assign to the content")
	out.Flags().StringP("file", "f", "", "Load message content from the given file if message not passed as a command line arg")
	out.Flags().StringP("data", "d", "", "Extra JSON data to save as part of the content")
	return out
}

func listMessagesCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "list TOPIC_ID",
		Short: "List messages",
		Long:  `List messages in a topic`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("A TOPIC_ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			path := fmt.Sprintf("/v1/topics/%s/messages", args[0])
			Client.Call("GET", path, nil, nil, nil)
		},
	}
	return out
}

func getMessagesCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "get MSG_ID [...MSG_IDS]",
		Short: "Get a message for one or more message IDs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("At least one message ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, mid := range args {
				path := fmt.Sprintf("/v1/messages/%s", mid)
				Client.Call("GET", path, nil, nil, nil)
			}
		},
	}
	return out
}

func deleteMessagesCommand() *cobra.Command {
	out := &cobra.Command{
		Use:   "delete MSG_ID [...MSG_IDS]",
		Short: "Delete a message by one or more message IDs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("At least one message ID must be provided")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, mid := range args {
				path := fmt.Sprintf("/v1/messages/%s", mid)
				Client.Call("DELETE", path, nil, nil, nil)
			}
		},
	}
	return out
}

func init() {
	rootCmd.AddCommand(messagesCommand())
}

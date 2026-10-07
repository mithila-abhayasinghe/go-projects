/*
Copyright © 2026 Mithila Abhayasinghe
*/
package cmd

import (
	"fmt"
	"github-activity/internal"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-activity <username>",
	Short: "Use GitHub API to fetch user activity and display it in the terminal",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("a single GitHub username is required (e.g. github-activity torvalds)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		uname := args[0]
		// resp := internal.FetchPublicRecordsByUser(uname)
		// fmt.Println(resp)
		internal.FetchPublicRecordsByUser(uname)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

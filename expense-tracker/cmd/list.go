/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		if verbose {
			fmt.Println("Verbose called")
		} else {
			fmt.Println("Verbose was not called")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

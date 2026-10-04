/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"expense-tracker/internal"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListExpense()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

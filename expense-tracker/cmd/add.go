/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "add expense description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "add the expense amount")

}

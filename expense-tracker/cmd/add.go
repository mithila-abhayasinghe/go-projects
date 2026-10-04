/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"expense-tracker/internal"
	"fmt"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := internal.AddExpense(description, amount)
		if err != nil {
			return err
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", exp.Id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "add expense description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "add the expense amount")

}

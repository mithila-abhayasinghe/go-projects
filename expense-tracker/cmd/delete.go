/*
Copyright © 2026 Mithila Abhayasinghe
*/
package cmd

import (
	"expense-tracker/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var didx int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an expense",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := internal.DeleteExpense(didx)
		if err != nil {
			return err
		}
		fmt.Printf("Expense deleted successfully (ID: %d)\n", exp.Id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&didx, "id", "i", 0, "delete an expense")
}

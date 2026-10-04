/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"expense-tracker/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// this package private
var (
	uidx  int
	udesc string
	uamt  float64
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an expense",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := internal.UpdateExpense(uidx, udesc, uamt)
		if err != nil {
			return err
		}
		fmt.Printf("Expense updated successfully (ID: %d)\n", exp.Id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&uidx, "id", "i", 0, "update an expense")
	updateCmd.Flags().StringVarP(&udesc, "description", "d", "", "updated description")
	updateCmd.Flags().Float64VarP(&uamt, "amount", "a", 0, "updated amount")

}

/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"expense-tracker/internal"

	"github.com/spf13/cobra"
)

var month int

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "shows the total expenses",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.SummaryExpense(month)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "summary of n month")
}

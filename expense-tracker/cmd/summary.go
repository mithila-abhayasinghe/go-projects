/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var month int

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "shows the total expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("summary called")
		fmt.Printf("Summary of month: %d\n", month)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "summary of n month")
}

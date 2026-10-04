/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// this package private
var uidx int

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		fmt.Printf("Update index: %d\n", uidx)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&uidx, "update", "u", 0, "update an expense")
}

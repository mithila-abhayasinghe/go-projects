/*
Copyright © 2026 Mithila Abhayasinghe
*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker [command] ",
	Short: "A lightweight CLI tool to track and summarize daily expenses",
	Long: `Expense Tracker is a fast, offline command-line utility to record
daily transactions, inspect logs in a clean tabular layout, and analyze
monthly spending totals.`,
	Example: `  # Add an expense
  expense-tracker add --description "Groceries" --amount 45.50

  # View all expenses
  expense-tracker list

  # Get total spending for August
  expense-tracker summary --month 8

  # Remove a mistake by ID
  expense-tracker delete --id 2`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var verbose bool

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// basic structures root.Type of flag.Flag Data type.
	// (pointer to variable, "long name", "shorthand", "default value", "description")
	// Flag for verbose output
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.expense-tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

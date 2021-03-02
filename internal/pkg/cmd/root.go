package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "nbp",
	Short: "NBP CLI tool",
	Long:  "CLI tool for easy access to National Bank of Poland API",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format, one of: json, table")
}

var (
	config       *Config
	outputFormat string
	rootCmd      = &cobra.Command{
		Use:   "nbp",
		Short: "NBP CLI tool",
		Long:  "CLI tool for easy access to National Bank of Poland API",
	}
)

type Config struct {
	Version   string
	Commit    string
	BuildDate string
}

// Execute executes the root command.
func Execute(cfg *Config) {
	config = cfg
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

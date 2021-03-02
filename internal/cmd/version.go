package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("NBP CLI v" + config.Version)
			fmt.Println()
			fmt.Println("commit: " + config.Commit)
			fmt.Println("built: " + config.BuildDate)
		},
	}
)

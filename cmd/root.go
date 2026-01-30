package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "recall",
	Short: "A minimal to-do list with a few amenities",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initStorage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

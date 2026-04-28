package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "recall",
	Short: "A minimal to-do list with a few amenities",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initStorage()
		applyDailyPenalty()
	},
}

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}

package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		_ = saveTasks([]Task{})
		color.Green("  Cleared all tasks!")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

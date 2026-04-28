package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clearCompletedCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clear completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		remaining := []Task{}

		for _, t := range tasks {
			if t.State < 2 || t.State == 3 {
				remaining = append(remaining, t)
			}
		}

		_ = saveTasks(remaining)
		color.Green(" Cleared completed tasks!")
	},
}

func init() {
	rootCmd.AddCommand(clearCompletedCmd)
}

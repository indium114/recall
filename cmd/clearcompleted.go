package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clearCompletedCmd = &cobra.Command{
	Use:   "clearcompleted",
	Short: "Clear completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		remaining := []Task{}

		for _, t := range tasks {
			if !t.State {
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

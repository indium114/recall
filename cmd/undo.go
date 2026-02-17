package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo <id>",
	Short: "Undo a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := atoi(args[0])
		tasks := loadTasks()

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].State = 0
				_ = saveTasks(tasks)

				penalty := t.Prio * 10
				decreaseXP(penalty)

				color.Yellow("󰓑 Lost %d XP\n", penalty)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}

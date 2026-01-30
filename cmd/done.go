package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Complete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := atoi(args[0])
		tasks := loadTasks()

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].State = true
				_ = saveTasks(tasks)

				reward := t.Prio * 10
				increaseXP(reward)

				color.Green("󱕣 Earned %d XP\n", reward)
				return
			}
		}

		color.Red(" Task does not exist")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

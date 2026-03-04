package cmd

import (
	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var busyCmd = &cobra.Command{
	Use:   "busy <id>",
	Short: "Set a task to busy",
	Run: func(cmd *cobra.Command, args []string) {
		id := atoi(args[0])
		tasks := loadTasks()

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].State = 1
				_ = saveTasks(tasks)

				color.Yellow("󰥔 Set task %d to busy", t.ID)
				return
			}
		}

		color.Red(" Task does not exist")
	},
}

func init() {
	rootCmd.AddCommand(busyCmd)
}

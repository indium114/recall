package cmd

import (
	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var backburnerCmd = &cobra.Command{
	Use:   "backburner <id>",
	Short: "Put a task on the back-burner",
	Run: func(cmd *cobra.Command, args []string) {
		id := atoi(args[0])
		tasks := loadTasks()

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].State = 3
				_ = saveTasks(tasks)

				color.Yellow("󰥔 Put task %d on the back-burner", t.ID)
				return
			}
		}

		color.Red(" Task does not exist")
	},
}

func init() {
	rootCmd.AddCommand(backburnerCmd)
}

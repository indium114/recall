package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo <id>",
	Short: "Undo a task (also works to unmark a task as busy or remove it from the back-burner)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := atoi(args[0])
		tasks := loadTasks()

		for i, t := range tasks {
			if t.ID == id {
				prevState := tasks[i].State

				tasks[i].State = 0
				_ = saveTasks(tasks)

				if prevState == 2 {
					penalty := t.Prio * 10
					decreaseXP(penalty)

					color.Yellow("󰓑 Lost %d XP\n", penalty)
				}
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}

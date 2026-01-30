package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()

		if len(tasks) == 0 {
			color.Green("󰄭 All tasks done!\n")
			return
		}

		color.White(" Tasks:\n")

		for _, t := range tasks {
			stat := " "
			if t.State {
				stat = "󰄲 "
			}

			line := fmt.Sprintf(
				"%d %s: %s (%d)\n",
				t.ID,
				stat,
				t.Name,
				t.Prio,
			)

			if t.Prio == 1 {
				color.Green(line)
			} else if t.Prio == 2 {
				color.Yellow(line)
			} else if t.Prio == 3 {
				color.Red(line)
			} else {
				color.White(line)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

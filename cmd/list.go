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

		var pending []Task
		var busy []Task
		var done []Task
		var backBurner []Task

		for _, t := range tasks {
			switch t.State {
			case 0:
				pending = append(pending, t)
			case 1:
				busy = append(busy, t)
			case 2:
				done = append(done, t)
			case 3:
				backBurner = append(backBurner, t)
			}
		}

		color.White(" Tasks:\n")

		if len(done) != 0 {
			color.White(" Completed")
		}

		for _, t := range done {
			stat := "󰄲 "

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

		if len(busy) != 0 {
			color.White(" Busy")
		}

		for _, t := range busy {
			stat := "󰥔 "

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

		if len(pending) != 0 {
			color.White(" Pending")
		}

		for _, t := range pending {
			stat := " "

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

		if len(backBurner) != 0 {
			color.White("󰀼 Back-Burner")
		}

		for _, t := range backBurner {
			stat := " "

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

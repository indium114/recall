package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <name> <priority>",
	Short: "Add a task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		prio := atoi(args[1])

		if prio < 1 || prio > 3 {
			color.Red(" Priority must be between 1 and 3")
			return
		}

		tasks := loadTasks()
		used := map[int]bool{}

		for _, t := range tasks {
			used[t.ID] = true
		}

		id := 0
		for used[id] {
			id++
		}

		tasks = append(tasks, Task{
			Name:  name,
			Prio:  prio,
			State: 0,
			ID:    id,
		})

		_ = saveTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

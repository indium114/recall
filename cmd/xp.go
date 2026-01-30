package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var xpCmd = &cobra.Command{
	Use:   "xp",
	Short: "Show XP",
	Run: func(cmd *cobra.Command, args []string) {
		xp := loadXP()
		color.Cyan(" XP: %d\n", xp.XP)
	},
}

func init() {
	rootCmd.AddCommand(xpCmd)
}

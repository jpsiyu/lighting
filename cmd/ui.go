package cmd

import (
	"github.com/jpsiyu/lighting/ui"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "user interface",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Run()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}

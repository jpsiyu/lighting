package cmd

import (
	"github.com/jpsiyu/lighting/ui"
	"github.com/jpsiyu/lighting/ui/demo"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "user interface",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Run()
	},
}

var uiDemoCmd = &cobra.Command{
	Use:   "demo",
	Short: "ui widgets demonstrate",
	Run: func(cmd *cobra.Command, args []string) {
		d := demo.NewDemo()
		d.Run()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)

	uiCmd.AddCommand(uiDemoCmd)
}

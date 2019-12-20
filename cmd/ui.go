package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "user interface",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ui called")
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}

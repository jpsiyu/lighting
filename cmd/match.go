package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "start match system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("match called")
	},
}

func init() {
	rootCmd.AddCommand(matchCmd)
}

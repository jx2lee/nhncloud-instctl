package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the versions command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print out the version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.2\n2021/03/20")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

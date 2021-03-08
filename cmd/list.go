package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/instance"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Check the list of instances.",
	Run: func(cmd *cobra.Command, args []string) {
		instance.ListInstances()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

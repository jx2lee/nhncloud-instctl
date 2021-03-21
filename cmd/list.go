package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Check the list of instances.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.InstanceListOutput(Region)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&Region, "region", "r", "", "set Cloud Region. (kr1, kr2, jp1)")
	listCmd.MarkFlagRequired("region")
}

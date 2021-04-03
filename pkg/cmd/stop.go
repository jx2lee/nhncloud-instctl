package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop <instance_name>",
	Short: "Stop the instance.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.PauseInstance(args[0], Region)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringVarP(&Region, "region", "r", "", "set Cloud Region. (kr1, kr2 or jp1)")
	stopCmd.MarkFlagRequired("region")

}

package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <instance_name>",
	Short: "\nStart the instance.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.StartInstance(args[0], Region)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&Region, "region", "r", "", "set Cloud Region. (kr1, kr2, jp1)")
	startCmd.MarkFlagRequired("region")
}

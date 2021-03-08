package cmd

import (
	"log"

	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the instance.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Not entered instance name. Please check and try again")
		}

		controller.PauseInstance(args[0])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

package cmd

import (
	"log"

	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the instance.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Not entered instance name. Please check and try again.")
		}

		controller.StartInstance(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

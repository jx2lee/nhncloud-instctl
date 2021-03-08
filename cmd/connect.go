package cmd

import (
	"log"

	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/spf13/cobra"
)

var name string
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to Instance.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Not entered instance name. Please check and try again.")
		}
		controller.SSHConnect(args[0])
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVar(&name, "name", "", "Enter the name of the instance you want to connect to.")
}

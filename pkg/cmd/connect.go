package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/controller"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to Instance.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logrus.Fatal("Not entered instance name. Please check and try again.")
		}
		controller.SSHConnect(args[0], Region)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVar(&name, "name", "", "Enter the name of the instance you want to connect to.")
	connectCmd.Flags().StringVarP(&Region, "region", "r", "", "set Cloud Region. (kr1, kr2 or jp1)")
	connectCmd.MarkFlagRequired("region")

}

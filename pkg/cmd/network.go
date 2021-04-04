package cmd

import (
	"github.com/jx2lee/nhncloud-instctl/pkg/network"

	"github.com/spf13/cobra"
)

var rdsCmd = &cobra.Command{
	Use:   "network",
	Short: "NHN Cloud Network Commands",
}

var getPortInfoCmd = &cobra.Command{
	Use:   "get-port",
	Short: "Look up the port on the iaas instance.",
	Run: func(cmd *cobra.Command, args []string) {
		network.GetIaasInstancePortInfo(IaasInstanceId)
	},
}

var attachFipCmd = &cobra.Command{
	Use:   "attach-fip",
	Short: "Connect the fip to the new instance.",
	Run: func(cmd *cobra.Command, args []string) {
		network.AttachFip(FIPId, PortId)
	},
}

func init() {
	rootCmd.AddCommand(rdsCmd)
	rdsCmd.AddCommand(getPortInfoCmd)
	rdsCmd.AddCommand(attachFipCmd)

	getPortInfoCmd.Flags().StringVarP(&IaasInstanceId, "iaas-id", "i", "", "set the iaas instance ID that connected the FIP.")
	getPortInfoCmd.MarkFlagRequired("iaas-id")

	attachFipCmd.Flags().StringVarP(&FIPId, "fip-id", "f", "", "set the floating_ip_id.")
	attachFipCmd.Flags().StringVarP(&PortId, "port-id", "p", "", "set the port_id.")
	attachFipCmd.MarkFlagRequired("fip-id")
	attachFipCmd.MarkFlagRequired("port-id")

}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Region string
var rootCmd = &cobra.Command{
	Use: "nhncloud-instctl",
	Version: "0.2",
	Short: `⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣄⣀⡀⠀⠀⠀⠀⣤⣤⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⡄
⣼⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⠀⠀⠀⠀⠀⠀⠀⢀⣴⣾⣿⣿⠿⠿⣿⣿⣿⣦⡀⠀⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡇
⣿⣿⡇⠀⢀⣀⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⠀⠀⣀⣠⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⠀⠀⠀⠀⠀⠀⢠⣾⣿⡿⠉⠀⠀⠀⠀⠈⠻⠟⠁⠀⣿⣿⡇⠀⠀⠀⣀⣤⣴⣶⣤⣄⡀⠀⠀⠀⣤⣤⡄⠀⠀⠀⢠⣤⣤⠀⠀⠀⣀⣤⣤⣦⣤⡀⣿⣿⡇
⣿⣿⡇⠀⠈⠻⢿⣿⣿⣶⣤⣀⠀⠀⠀⠀⠀⣿⣿⡇⠀⠀⢀⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⠀⠀⢸⣿⣿⠀⠀⠙⠿⣿⣿⣷⣦⣄⡀⠀⠀⠀⠀⢸⣿⣿⠀⠀⠀⠀⠀⠀⣼⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡇⠀⢠⣾⣿⠿⠛⠛⠻⣿⣿⣆⠀⠀⣿⣿⡇⠀⠀⠀⢸⣿⣿⠀⢀⣾⣿⡿⠛⠛⠛⢿⣿⣿⡇
⣿⣿⡇⠀⠀⠀⠀⠈⠙⠻⣿⣿⣿⣶⣄⠀⠀⣿⣿⡇⠀⠀⠛⠻⠿⠿⠿⠿⠿⠿⠿⠟⠋⠀⠀⢸⣿⣿⠀⠀⠀⠀⠀⠉⠛⠿⣿⣿⣷⣦⡀⠀⢸⣿⣿⠀⠀⠀⠀⠀⠀⢿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡇⠀⣾⣿⡏⠀⠀⠀⠀⠈⣿⣿⡆⠀⣿⣿⡇⠀⠀⠀⢸⣿⣿⠀⢸⣿⡟⠀⠀⠀⠀⠀⢻⣿⡇
⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠛⠉⠀⠀⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠋⠁⠀⢸⣿⣿⠀⠀⠀⠀⠀⠀⠘⣿⣿⣧⡀⠀⠀⠀⠀⢀⣠⣤⡀⠀⣿⣿⡇⠀⢿⣿⣇⠀⠀⠀⠀⢠⣿⣿⠃⠀⣿⣿⡇⠀⠀⠀⣸⣿⡟⠀⢸⣿⣷⡀⠀⠀⠀⢀⣾⣿⡇
⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡿⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣷⣶⣶⣾⣿⣿⠟⠁⠀⣿⣿⡇⠀⠈⠻⣿⣷⣶⣶⣶⣿⡿⠋⠀⠀⠹⣿⣿⣶⣶⣾⣿⡿⠃⠀⠀⠻⣿⣿⣶⣤⣶⣿⣿⣿⡇
⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠛⠋⠉⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠈⠉⠛⠛⠋⠉⠀⠀⠀⠀⠀⠈⠉⠛⠛⠋⠉⠀⠀⠀⠀⠀⠈⠉⠛⠛⠉⠁⠉⠉⠁

Instance CLI for NHN Cloud
CLI for easy management of instances created by NHN Cloud.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

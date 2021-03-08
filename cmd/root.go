package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "nhncloud-instctl",
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

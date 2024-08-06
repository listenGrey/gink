package cmd

import (
	"fmt"
	"gink/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ip] [port]",
	Short: "Add a new destination",
	Long:  `Add a new IP address and port as a destination for sending files.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		config.AppConfig.Destinations[0] = ip
		if err := config.SaveConfig(); err != nil {
			fmt.Printf("Error saving configuration: %s\n", err)
		} else {
			fmt.Println("Destination added.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

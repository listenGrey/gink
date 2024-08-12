package cmd

import (
	"fmt"
	"gink/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ip]",
	Short: "Add a new destination",
	Long:  `Add a new IP address and port as a destination for sending files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		err := config.AddDestination(ip)
		if err != nil {
			fmt.Printf("Error saving configuration: %s\n", err)
		} else {
			fmt.Println("Destination added.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

package cmd

import (
	"fmt"
	"gink/config"
	"github.com/spf13/cobra"
)

var desCmd = &cobra.Command{
	Use:   "des",
	Short: "List all destinations",
	Long:  `List all configured destinations with their corresponding numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configured destinations:")
		for k, ip := range config.AppConfig.Destinations {
			fmt.Printf("%s..........%d\n", ip, k+1)
		}
	},
}

func init() {
	rootCmd.AddCommand(desCmd)
}

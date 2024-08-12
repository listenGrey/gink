package cmd

import (
	"fmt"
	"gink/config"
	"github.com/spf13/cobra"
)

var protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "List all protocols",
	Long:  `List all configured protocols.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Using %s protocol to transport files\n", config.AppConfig.Protocols[0])
		fmt.Printf("There are protocols you can choose, change it on config file:\n")
		for _, protocol := range config.AppConfig.Protocols {
			fmt.Println(protocol)
		}
	},
}

func init() {
	rootCmd.AddCommand(protocolCmd)
}

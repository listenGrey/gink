package cmd

import (
	"fmt"
	"gink/config"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local [path]",
	Short: "Set the local directory to save received files",
	Long:  `Set the local directory where the received files will be saved.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.AppConfig.LocalDirection = args[0]
		if err := config.SaveConfig(); err != nil {
			fmt.Printf("Error saving configuration: %s\n", err)
		} else {
			fmt.Println("Local save path updated.")
		}
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}

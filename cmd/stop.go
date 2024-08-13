package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the application",
	Long:  `This command stops the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application is stopping...")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

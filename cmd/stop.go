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
		os.Exit(0) // 此处简单示例，实际上可能需要优雅地关闭资源
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

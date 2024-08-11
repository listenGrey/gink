package cmd

import (
	"fmt"
	"gink/pkg/transfer"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Long:  `This command starts the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application is running...")
		var Trans transfer.Transfer
		Trans = &transfer.TCPTransfer{} // 使用TCP协议
		err := Trans.Receive()          // 启动监听服务
		if err != nil {
			fmt.Printf("Failed to receive file: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

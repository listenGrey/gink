package cmd

import (
	"fmt"
	"gink/config"
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
		switch config.AppConfig.Protocols[0] {
		case "websocket":
			Trans = &transfer.WebSocketTransfer{} // 使用websocket协议
		case "tcp":
			Trans = &transfer.TCPTransfer{} // 使用TCP协议
		default:
			fmt.Println("Protocol error")
		}
		err := Trans.Receive() // 启动监听服务
		if err != nil {
			fmt.Printf("Failed to receive file: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

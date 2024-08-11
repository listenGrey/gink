package cmd

import (
	"fmt"
	"gink/pkg/transfer"
	"github.com/spf13/cobra"
)

var (
	filepath         string
	destinationIndex string
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a file",
	Long:  `Send a file to a specified destination.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Trans transfer.Transfer
		Trans = &transfer.TCPTransfer{} // 使用TCP协议
		err := Trans.Send(filepath, destinationIndex)
		if err != nil {
			fmt.Printf("Failed to send file: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "Path to the file to send")
	sendCmd.Flags().StringVarP(&destinationIndex, "destination", "d", "", "Destination IP:Port")
	sendCmd.MarkFlagRequired("filepath")
	sendCmd.MarkFlagRequired("destination")
}

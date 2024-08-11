package cmd

import (
	"fmt"
	"gink/pkg/transfer"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Displays the history of file transfers",
	Long:  `This command displays a detailed history of all file transfers.`,
	Run: func(cmd *cobra.Command, args []string) {
		history := transfer.GetHistory()
		for k, record := range history {
			if k == 0 {
				continue
			}
			if record.Receive {
				if record.Success {
					fmt.Printf("Receive from %s at %s, Name: %s, Success: %t\n", record.Destination, record.Time, record.FileName, record.Success)
				} else {
					fmt.Printf("Receive from %s at %s, Name: %s, Success: %t, Error: %s\n", record.Destination, record.Time, record.FileName, record.Success, record.ErrorMessage)
				}
			} else {
				if record.Success {
					fmt.Printf("Send to %s at %s, Name: %s, Success: %t\n", record.Destination, record.Time, record.FileName, record.Success)
				} else {
					fmt.Printf("Send to %s at %s, Name: %s, Success: %t, Error: %s\n", record.Destination, record.Time, record.FileName, record.Success, record.ErrorMessage)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}

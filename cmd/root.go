package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gink",
	Short: "Gink is a P2P file transfer system",
	Long:  `Gink is a CLI application designed for peer-to-peer file transfers.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you can define flags and configuration settings.
}

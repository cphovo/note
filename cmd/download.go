package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download default data if not available",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: download gojieba dataset if not available
		fmt.Println("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

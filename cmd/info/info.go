package info

import (
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "System information commands",
	Long:  `Commands related to retrieving system information like disk usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	addDiskUsageCommand()
}

func addDiskUsageCommand() {
	InfoCmd.AddCommand(diskUsageCmd)
}

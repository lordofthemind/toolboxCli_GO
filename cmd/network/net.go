package network

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "network",
	Short: "Network-related commands",
	Long:  `Network-related commands like pinging a remote URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	addPingCommand()
}

func addPingCommand() {
	NetCmd.AddCommand(pingCmd)
}

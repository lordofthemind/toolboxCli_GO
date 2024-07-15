package cmd

import (
	"os"

	"github.com/lordofthemind/toolboxCli_GO/cmd/info"
	"github.com/lordofthemind/toolboxCli_GO/cmd/network"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toolboxCli_GO",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(network.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}

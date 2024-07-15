/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"

	"github.com/spf13/cobra"
)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskUsage called")
		usage := du.NewDiskUsage(".")
		fmt.Printf("Free Disk Space: %d bytes\n", usage.Free())
		fmt.Printf("Used Disk Space: %d bytes\n", usage.Size()-usage.Free())
		fmt.Printf("Total Disk Space: %d bytes\n", usage.Size())

	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Display disk usage information",
	Long:  `Display detailed information about disk usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskUsage called")
		usage := du.NewDiskUsage(".")
		fmt.Printf("Free Disk Space: %d bytes\n", usage.Free())
		fmt.Printf("Used Disk Space: %d bytes\n", usage.Size()-usage.Free())
		fmt.Printf("Total Disk Space: %d bytes\n", usage.Size())
	},
}

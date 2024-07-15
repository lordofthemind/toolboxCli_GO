package network

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = &http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := "https://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping a remote URL",
	Long:  `Ping a remote URL and get the status code.`,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Response: ", resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to ping")
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
}

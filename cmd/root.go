package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

const (
	URL_FLAG = "url"
)

var rootCmd = &cobra.Command{
	Use:   "web-crawler",
	Short: "A web crawler that discovers and traverses links starting from a specified URL.",
	Long: `A web crawler tool that:
- Starts from a specified root URL
- Discovers and follows links using BFS
- Respects robots.txt and politeness policies
- Outputs crawl results in structured formats (JSON/text)
- Supports concurrency and domain limitation

Example usage:
  web-crawler -url https://example.com -depth 3 -output sitemap.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		targetUrl, err := cmd.Flags().GetString(URL_FLAG)

		if err != nil {
			return err
		}

		err = IsValidUrl(targetUrl)
		if err != nil {
			return err
		}

		fmt.Printf("URL: %s", targetUrl)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP(URL_FLAG, "u", "", "A URL to be crawled")
	if err := rootCmd.MarkFlagRequired(URL_FLAG); err != nil {
		panic(err)
	}
}

func IsValidUrl(str string) error {
	u, err := url.ParseRequestURI(str)

	if err != nil {
		return err
	}
	if u.Scheme == "" || u.Host == "" {
		return errors.New("Invalid url")
	}

	return nil
}

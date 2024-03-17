package cmd

import (
	"fmt"
	"github.com/straski/go-wsdl/downloader"
	"os"

	"github.com/spf13/cobra"
)

var url string
var targetDir string

var rootCmd = &cobra.Command{
	Use:   "go-wsdl",
	Short: "A website download tool",
	Long:  `Download a whole website by providing the site's URL.`,
	Run: func(cmd *cobra.Command, args []string) {

		result := downloader.Download(url, targetDir)

		fmt.Printf("- %d pages\n", result.Ahrefs)
		fmt.Printf("- %d scripts\n", result.Scripts)
		fmt.Printf("- %d images\n", result.Images)
		fmt.Printf("- %d link resources\n", result.Links)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "http://books.toscrape.com/index.html", "Website URL")
	rootCmd.Flags().StringVarP(&targetDir, "dir", "d", "output", "Directory to save the files to.")
}

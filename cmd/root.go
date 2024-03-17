package cmd

import (
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
		// call downloader.download
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

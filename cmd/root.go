package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/straski/go-wsdl/downloader"
	"log"
	"net/http"
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

		blue := color.FgLightBlue.Render
		color.Warnf("\n- Download %s ", blue(url))
		color.Warnf("to directory '%s'\n", blue(targetDir))
		color.Warn.Println("> Downloading...")

		result := downloader.Download(url, targetDir)

		color.Info.Println("> Done. Downloaded:")
		fmt.Printf("- %d pages\n", result.Ahrefs)
		fmt.Printf("- %d scripts\n", result.Scripts)
		fmt.Printf("- %d images\n", result.Images)
		fmt.Printf("- %d link resources\n", result.Links)

		color.Warn.Println("> Scanning CSS files for additional resources...")

		if css, err := downloader.ScanCss(targetDir, url); err != nil {
			color.Error.Println("_ Error downloading additional resources")
		} else {
			for _, file := range css {
				fmt.Printf("_ Found file %s\n", file)
			}
			fmt.Printf("- Downloaded %d additional resources.\n", len(css))
		}

		color.Warn.Print("> Starting web server...")
		color.Warn.Printf(blue("\n- Open http://localhost:1337 or %s/index.html in your browser.\n"), targetDir)
		color.Warn.Println("- Press CTRL+C to stop the web server.\n")

		http.Handle("/", http.FileServer(http.Dir(targetDir)))
		if err := http.ListenAndServe(":1337", nil); err != nil {
			color.Error.Println(
				"Couldn't start web server - access the page by opening index.html in the output directory.")
			log.Fatal("ListenAnServe:", err)
		}
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

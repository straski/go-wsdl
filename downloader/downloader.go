package downloader

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"path/filepath"
	"slices"
	"time"
)

// Result holds all the resource links
type Result struct {
	Links   int
	Ahrefs  int
	Scripts int
	Images  int
}

func newResult(links, ahrefs, scripts, images int) *Result {
	return &Result{Links: links, Ahrefs: ahrefs, Scripts: scripts, Images: images}
}

// Download downloads the website resources
func Download(url, targetDir string) *Result {
	defer timer("Download")()
	var links, scripts, images, ahrefs []string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("filepath", targetDir+(r.URL.Path))
		r.Ctx.Put("dirname", targetDir+filepath.Dir(r.URL.Path))
		fmt.Printf("_ %s\n", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		ensureDirectoryExists(r.Ctx.Get("dirname"))
		r.Save(r.Ctx.Get("filepath"))

	})

	c.OnHTML("link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !slices.Contains(links, link) {
			links = append(links, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !slices.Contains(ahrefs, link) {
			ahrefs = append(ahrefs, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if !slices.Contains(images, link) {
			images = append(images, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if !slices.Contains(scripts, link) {
			scripts = append(scripts, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(url)

	return newResult(len(links), len(ahrefs), len(scripts), len(images))
}

// ensureDirectoryExists check if dir exists and creates it if not
func ensureDirectoryExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return
		}
	}
}

// timer is a helper for measuring execution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("_ %s took %v\n", name, time.Since(start))
	}
}

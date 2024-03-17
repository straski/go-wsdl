package downloader

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly/v2"
	"golang.org/x/exp/slices"
	"io"
	"io/fs"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
	"strings"
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

	c := colly.NewCollector(
		colly.MaxDepth(3),
		colly.Async(),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 8})

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
	c.Wait()

	return newResult(len(links), len(ahrefs), len(scripts), len(images))
}

// ScanCss scans CSS files for additional resources and downloads them
func ScanCss(targetDir, url string) (res []string, err error) {
	defer timer("Scan")()
	pUrl, _ := url2.Parse(url)

	for _, cssFile := range findFileByExtension(targetDir, ".css") {
		for _, typ := range []string{".eot", ".woff", ".svg", ".ttf"} {
			if fontResources, err := getFonts(cssFile, typ, filepath.Dir(cssFile), targetDir, pUrl.Scheme+"://"+pUrl.Host); err != nil {
				return res, err
			} else {
				res = append(res, fontResources...)
			}
		}
	}
	return res, nil
}

// getFonts searches for font links in CSS files
func getFonts(cssFile, typ, parentDir, targetDir, url string) (res []string, err error) {
	for _, fontResourceLine := range findStringInFile(cssFile, typ) {
		fontResourcesSplit := strings.Split(fontResourceLine, " ")
		for _, fontResource := range fontResourcesSplit {
			if strings.Contains(fontResource, typ) {
				fontResourcePath := getStringBetweenStrings(fontResource, "url('", "')")
				if fontResourcePath != "" {
					fullFontPath := parentDir + "/" + fontResourcePath
					fontDir := filepath.Dir(fullFontPath)
					fullFontPathAbs := filepath.Dir(fullFontPath) + "/" + filepath.Base(fullFontPath)
					fullFontUrl := strings.Replace(fullFontPath, targetDir, url, 1)

					ensureDirectoryExists(fontDir)

					if err := downloadUrlToFile(fullFontUrl, fullFontPathAbs); err != nil {
						return res, err
					}

					res = append(res, fullFontUrl)
				}
			}
		}
	}
	return res, err
}

// findFileByExtension finds a file by extension name
func findFileByExtension(dir, ext string) []string {
	var a []string
	filepath.WalkDir(dir, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

// findStringInFile finds a string in a given file
func findStringInFile(filepath, needle string) (res []string) {
	f, err := os.Open(filepath)
	if err != nil {
		return res
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), needle) {
			res = append(res, scanner.Text())
		}

		line++
	}
	return res
}

// getStringBetweenStrings gets a string between to string delimiters
func getStringBetweenStrings(haystack, start, end string) (result string) {
	s := strings.Index(haystack, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(haystack[s:], end)
	if e == -1 {
		return
	}
	return haystack[s : s+e]
}

func downloadUrlToFile(url, filepath string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
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

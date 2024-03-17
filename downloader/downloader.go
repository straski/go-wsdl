package downloader

import (
	"fmt"
	"time"
)

func Download(url, targetDir string) {
	defer timer("Download")()
}

// timer is a helper for measuring execution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("_ %s took %v\n", name, time.Since(start))
	}
}

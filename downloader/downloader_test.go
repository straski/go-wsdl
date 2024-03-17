package downloader

import (
	"os"
	"testing"
)

func TestDownloadUrlToFile(t *testing.T) {
	testFile := "../_testrun/TestDownloadUrlToFile.jpg"
	testUrl := "http://books.toscrape.com/media/cache/fe/72/fe72f0532301ec28892ae79a629a293c.jpg"

	err := downloadUrlToFile(testUrl, testFile)
	if err != nil {
		t.Fatalf(`Test failed %v`, err)
	}

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf(`Test failed %v`, err)
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	testDir := "../_testrun/TestEnsureDirectoryExists"
	ensureDirectoryExists(testDir)

	_, err := os.Stat(testDir)
	if os.IsNotExist(err) {
		t.Fatalf(`Test failed, directory %s does not exist (%v).`, testDir, err)
	}
}

func TestGetStringBetweenStrings(t *testing.T) {
	expected := "TestGetStringBetweenStrings"
	testString := "url('TestGetStringBetweenStrings')"
	testStart := "url('"
	testEnd := "')"

	r := getStringBetweenStrings(testString, testStart, testEnd)

	if r != expected {
		t.Fatalf(`Test failed, got %s, expected %s`, r, expected)
	}
}

func TestFindStringInFile(t *testing.T) {
	stringToFind := "&colly.LimitRule"
	fileToLookIn := "downloader.go"
	testLine := "\terr := c.Limit(&colly.LimitRule{DomainGlob: \"*\", Parallelism: 8})"

	matchingLines := findStringInFile(fileToLookIn, stringToFind)

	for _, i := range matchingLines {
		if i != testLine {
			t.Fatalf(`Test failed, got %s, expected %s`, matchingLines, testLine)
		}
	}
}

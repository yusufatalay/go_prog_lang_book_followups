package main

import (
	"fmt"
	"io"
	"links"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Crawl the web breadth-first
	// starting from the command-line arguments
	breadthFirst(crawl, os.Args[1:])

}

// breadthFirst calls f for each item in the worklist
// Any items returned by f are addded to the worklist
// f is called at most once for each item
func breadthFirst(f func(item string) []string, worklist []string) {

	// have ever visited that links ? Prevent the cyclicity
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// that "..." is for appending everything that returned by the
				// f to the worklist
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("cannot GET %s: %v", url, err)
	}
	// ignoring the error
	doc, err := io.ReadAll(resp.Body)
	// prepare the url to splitting for folder creation proccess
	trimmedurl := ""
	if strings.HasPrefix(url, "https://") {
		trimmedurl = strings.TrimPrefix(url, "https://")
	} else if strings.HasPrefix(url, "http://") {
		trimmedurl = strings.TrimPrefix(url, "http://")
	}
	temparr := strings.Split(trimmedurl, "/")
	filename := temparr[len(temparr)-1]
	filename = strings.TrimSuffix(filename, ":") + ".html"

	// no need to create nested folders by hand, os.MkdirAll will handle that.
	filelink := "./CrawlResult/" + trimmedurl
	os.MkdirAll(filelink, os.ModePerm)
	filepath := filelink + "/" + filename
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("cannot create file %s", filepath)
	}
	file.WriteString(string(doc))
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()
	return list
}

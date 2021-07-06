package main

import (
	"fmt"
	"links"
	"log"
	"os"
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
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

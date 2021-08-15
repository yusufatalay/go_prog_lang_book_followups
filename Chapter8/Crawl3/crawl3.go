// Contains exercise 8.6
package main

import (
	"flag"
	"fmt"
	"links"
	"log"
)

var depth = flag.Int("depth", 3, "URLs reachable by at most -depth links will be fetched")

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	flag.Parse()
	worklist := make(chan []string)  // lists of URLs , may have duplicates
	unseenLinks := make(chan string) //de-duplicated list

	// Add command-line arguments to worklist
	go func() { worklist <- flag.Args() }()

	// Create 20 crawler goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				founlinks := crawl(link)
				go func() { worklist <- founlinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers
	seen := make(map[string]bool)
	// depthList is a map to denote how many occurentce a URL has
	depthList := make(map[string]int)

	for list := range worklist {
		for _, link := range list {
			depthList[link]++
			if !seen[link] && depthList[link] < *depth {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

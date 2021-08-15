package main

import (
	"fmt"
	"links"
	"log"
	"os"
)

// tokens is a counting semaphore  used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	worklist := make(chan []string)

	var n int // number of pending sends to worklist

	// Start with the command line args

	n++
	go func() { worklist <- os.Args[1:] }()

	// crawl the web concurrently

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// this global array will hold the html tags that contains text

func main() {
	w, i, err := CountWordsAndImages(os.Args[1])

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%s contains:\n\t%d words\n\t%d images\n", os.Args[1], w, i)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {

	nodes := make([]*html.Node, 0) // hold all unvisited nodes

	nodes = append(nodes, n)

	for len(nodes) > 0 {
		n = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		if n.Type == html.TextNode {
			words += wordCounter(n.Data)
		}

		// add connected nodes to the list
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}
	}
	return

}

// wordCounter count the white space seperated wors and returns the result
func wordCounter(s string) (count int) {
	wordlist := strings.Split(s, " ")
	count = len(wordlist)
	return
}

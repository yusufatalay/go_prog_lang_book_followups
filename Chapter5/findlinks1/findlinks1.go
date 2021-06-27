// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println("2")
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println("1")
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		fmt.Println("3")

		return links
	}
	if n.FirstChild == nil {
		fmt.Println("6")
		return links
	}
	if n.FirstChild.NextSibling == nil {
		fmt.Println("7")
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println("4")
				links = append(links, a.Val)
			}
		}
	}

	//	links = visit(links, n.NextSibling)

	fmt.Println("5")
	c := n.FirstChild
	links = visit(links, c)
	c = c.NextSibling

	//	for c := n.FirstChild; c != nil; c = c.NextSibling {
	//		links = visit(links, c)
	//	}
	return links

}

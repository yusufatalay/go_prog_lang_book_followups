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
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	//	findEverything(doc)
}

// visit appends to links each link found in n and result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "script" || n.Data == "image" {
			for _, src := range n.Attr {
				if src.Key == "src" {
					links = append(links, src.Val)
				}
			}
		}
		if n.Data == "link" {
			for _, l := range n.Attr {
				if l.Key == "href" {
					links = append(links, l.Val)
				}
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

// findEverything prints all text nodes in html doc except <style> and <scritp>
func findEverything(n *html.Node) {

	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, tag := range n.Attr {
			fmt.Printf("<%s>%s</%s>\n", n.Data, tag.Val, n.Data)
		}
	}
	if n.FirstChild != nil {
		findEverything(n.FirstChild)
	}
	if n.NextSibling != nil {
		findEverything(n.NextSibling)
	}
}

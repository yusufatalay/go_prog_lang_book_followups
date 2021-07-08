package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalf("http.Get error: %v\n", err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html.Parse() error: %v\n", err)
	}
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)
	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild == nil {
				fmt.Printf("%*s<%s %s", depth*2, "", n.Data, getAttributes(n))
			} else {
				fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, getAttributes(n))
			}
			depth++
		}
		if n.Type == html.TextNode {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
			depth++
		}
		if n.Type == html.CommentNode {
			fmt.Printf("<!--%s-->\n", n.Data)
			depth++
		}

	}
	endElement = func(n *html.Node) {
		if n.Type == html.TextNode || n.Type == html.CommentNode {
			depth--
		}
		if n.Type == html.ElementNode {
			depth--
			if n.FirstChild == nil {
				fmt.Printf("%s/>\n", "")
			} else {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}

	}

	forEachNode(doc, startElement, endElement)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// getAttributes builds a string that consists of the given node's attributes
// and returns it
func getAttributes(n *html.Node) string {
	attrs := ""
	for _, atr := range n.Attr {
		attrs += fmt.Sprintf(" %s=\"%s\" ", atr.Key, atr.Val)
	}
	return attrs
}

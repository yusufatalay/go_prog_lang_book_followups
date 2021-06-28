package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline : %s", err)
		os.Exit(1)
	}
	outline(nil, doc)
	res := make(map[string]int)
	res = attramount(res, doc)
	for key, val := range res {
		fmt.Printf("Amount of attrbiute  %s : %d \n", key, val)
	}
}

func outline(stack []string, n *html.Node) {

	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// attramount populates a mapping from element names to the number of elements with that name in given html doc
func attramount(m map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		m[n.Data] += 1
	}

	if n.FirstChild != nil {
		attramount(m, n.FirstChild)
	}
	if n.NextSibling != nil {
		attramount(m, n.NextSibling)
	}
	return m
}

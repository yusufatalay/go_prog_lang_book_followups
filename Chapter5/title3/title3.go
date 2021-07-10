// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Printf("get %s :  %v\n", os.Args[1], err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		fmt.Printf("error while parsing %s --> %v\n", os.Args[1], err)
	}
	title, err := soleTitle(doc)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(title)

}

func soleTitle(doc *html.Node) (title string, err error) {

	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic ; carry on panicking
		}
	}()
	// Bail out of recurssion if we find more then one non-empty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

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

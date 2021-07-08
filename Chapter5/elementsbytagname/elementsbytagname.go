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
		fmt.Printf("cannot get %s : %v\n", os.Args[1], err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)

	// example part
	images := ElmentsByTagName(doc, os.Args[2:]...)

	for _, e := range images {
		fmt.Println(e.Data)
	}
	// ------
}

func ElmentsByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node

	for _, i := range name {
		if doc.Data == i {
			result = append(result, doc)
		}
	}

	if doc.FirstChild != nil {
		result = append(result, ElmentsByTagName(doc.FirstChild, name...)...)
	}
	if doc.NextSibling != nil {
		result = append(result, ElmentsByTagName(doc.NextSibling, name...)...)
	}
	return result
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findlink makes a get request to given url parses the response for more links and returns them
func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	// defer this func instead of calling it before every return
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, err
	}

	return visit(nil, doc), nil
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

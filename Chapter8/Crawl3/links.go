// Package links provides a link-extraction function
package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// Extract makes and HTTP GET request to the specified URL, parses the response
// as HTML, and returns the links in the HTML document.
func Extract(url string, cancel <-chan struct{}) ([]string, error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	// just don't need it anymore
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	// anon function
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// we no longer take the raw href value instead we parse it as
				// URL. So the link is in absolute form and ready to use in a
				// http.Get function.
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	// we no longer need the post function.
	forEachNode(doc, visitNode, nil)
	return links, nil
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

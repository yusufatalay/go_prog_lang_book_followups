// Contains exercise 8.11
package main

import (
	"fmt"
	"net/http"
	"os"
)

// request performs a http get request to given url
// and returns the given URL.
func request(url string, cancel <-chan struct{}) string {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = cancel
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return resp.Request.URL.String()

}

func getFirstResp(urls []string) string {

	cancel := make(chan struct{})
	respChan := make(chan string, len(urls))
	for _, url := range urls {
		go func(singleURL string) { respChan <- request(singleURL, cancel) }(url)
	}
	defer func() { close(cancel) }()
	return <-respChan
}

func main() {
	fastest := getFirstResp(os.Args[1:])
	fmt.Println(fastest)
}

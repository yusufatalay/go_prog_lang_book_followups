// Fetch downloads the URL and rturns the name and lengthof the local file
package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fetch(os.Args[1])
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	//	acts the same way with the commented code
	defer func() { err = f.Close() }()

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from copy, if any.
	//	if closeErr := f.Close(); err == nil {
	//		err = closeErr
	//	}
	return local, n, err
}

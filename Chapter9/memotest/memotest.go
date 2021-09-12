package main

import (
	"http"
	"io/ioutil"
	"memo1"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := memo1.New(httpGetBody)
}

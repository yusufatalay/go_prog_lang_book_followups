package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Comic struct {
	Index      int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

const baseURL = "http://xkcd.com/"
const urlpostfix = "/info.0.json"
const fileext = ".json"

// fileExists checks if given file exists adn returns accordingly
func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getComic(index int) (*Comic, error) {

	fmt.Println(string(index))
	resp, err := http.Get(baseURL + strconv.Itoa(index) + urlpostfix)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		log.Fatalf("\n[!]That was all of the comics\n")
		resp.Body.Close()
		errorTheEnd := errors.New("No more comics left to search")
		return nil, errorTheEnd
	}

	var comic Comic

	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &comic, nil

}

func main() {
	index := 1
	for {
		comic, err := getComic(index)

		if err != nil {
			log.Fatalf("%v", err)
			break
		}

		fmt.Printf("%s/%s/%s\n", comic.Day, comic.Month, comic.Year)
		fmt.Printf("%s\n", comic.Title)
		fmt.Printf("%s\n", comic.Transcript)
		fmt.Println("--------------------------------------------------")
		index++
	}
}

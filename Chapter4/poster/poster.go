package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// MoviePoster struct will only be used for holding the poster URL this way the program will look more tidier

type MoviePoster struct {
	Title     string `json:Title`
	PosterURL string `json:"Poster"`
}

// getPosterURL returns URL of the given movie's poster
func getPosterURL(movieURL string) *MoviePoster {
	resp, err := http.Get(movieURL)

	if err != nil {
		fmt.Errorf("Error occured while requesting for the movie poster --> %v\n", err)
		return nil
	}
	defer resp.Body.Close()
	var poster MoviePoster

	if err := json.NewDecoder(resp.Body).Decode(&poster); err != nil {
		fmt.Errorf("Error occured while parsing the JSON --> %v\n", err)
		return nil
	}

	return &poster
}

func downloadAndOpen(poster *MoviePoster) {
	title := poster.Title + ".jpg"
	posterURL := poster.PosterURL

	// downloading the poster
	wgetcmd := exec.Command("wget", "-O"+title, posterURL)

	err := wgetcmd.Run()

	if err != nil {
		fmt.Errorf("Error occured while downloading the poster --%v\n", err)
		return
	}

	opencmd := exec.Command("xdg-open", title)

	err = opencmd.Run()

	if err != nil {
		fmt.Errorf("Error occured while displaying the poster --> %v\n", err)
		return
	}
	return

}

func main() {
	apikey_byte, err := ioutil.ReadFile("apikey.txt")
	// turn apikey to string from byte slice
	apikey := string(apikey_byte)
	// get rid of the newline character at the end of the apikey

	apikey = strings.TrimSuffix(apikey, "\n")
	if err != nil {
		fmt.Errorf("Error occured while reading the apikey.txt file provide your omdbapi apikey within a text file called apikey.txt --> %v\n", err)
	}

	movietitle := os.Args[1:]

	movieURL := "http://www.omdbapi.com/?apikey=" + string(apikey) + "&t=" + movietitle
	poster := getPosterURL(movieURL)

	downloadAndOpen(poster)
}

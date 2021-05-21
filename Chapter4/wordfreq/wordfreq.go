package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	wordfreq := make(map[string]int)

	if err != nil {
		fmt.Println("Cannot read ", err)
	}

	inputText := string(input)
	wordlist := strings.Split(inputText, " ")

	for _, v := range wordlist {
		wordfreq[v]++
	}

	for k, v := range wordfreq {
		fmt.Printf("%s\t%-20d\n", k, v)
	}

}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := []string(os.Args[1:])
	// input : go is a nice language
	fmt.Println(joinvariadic(":", input...))
	// output: go:is:a:nice:language
}

func joinvariadic(seperator string, verbs ...string) string {
	var result string
	for _, v := range verbs {
		result += v + seperator
	}
	// remove the trailing seperator
	result = strings.TrimSuffix(result, seperator)
	return result
}

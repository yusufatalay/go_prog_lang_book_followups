package main

import (
	"bytes"
	"fmt"
	"strings"
)

// comma , inserts comma after every 3 integers

func comma(s string) string {
	n := len(s)
	mod3 := n % 3

	var rString bytes.Buffer

	if mod3 == 0 {
		mod3 = 3 // if it has multiple of 3's digits the we can divide it into triple groups
	}
	// get the first part
	rString.WriteString(s[:mod3])

	for i := mod3; i+3 < n; i += 3 {
		rString.WriteByte(',')
		rString.WriteString(s[i : i+3])
	}

	return rString.String()
}

func handleComma(s string) string {

	dotIndex := strings.LastIndex(s, ".")

	if dotIndex == -1 {
		return comma(s)
	}
	afterDot := s[dotIndex:]
	beforeDot := s[:dotIndex]
	beforeDot = comma(s[:dotIndex])
	return beforeDot + afterDot

}

func main() {

	ex := "723741327132727183721573.82512358"

	fmt.Printf(handleComma(ex) + "\n")

}

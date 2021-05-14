package main

import "fmt"

// comma , inserts comma after every 3 integers

func comma(s string) string {
	n := len(s)

	r := ""

	if n <= 3 {

		return s
	}

	for i := len(s); i >= 3; i -= 3 {
		r = s[i-3:i] + "," + r
	}
	return r
}

func main() {

	ex := "1234567"

	fmt.Printf(comma(ex) + "\n")

}

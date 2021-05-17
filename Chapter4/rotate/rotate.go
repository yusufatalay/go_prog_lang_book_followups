package main

import "fmt"

func rev(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, r int, d string) {

	switch d {
	case "l":
		rev(s[:r])
		rev(s[r:])
		rev(s[:])
	case "r":
		rev(s[:])
		rev(s[:r])
		rev(s[r:])

	}

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotate(a, 2, "l")
	fmt.Println(a)
}

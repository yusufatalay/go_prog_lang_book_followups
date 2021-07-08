// max and min are variadic functions. They do as their name implies
package main

import (
	"fmt"
	"sort"
)

func main() {

	input := []int{3, 4, 6, 2, 6, 2, 4456, 367, 3, 67, 2, 8, 347, 45, 4623, 534, 623, 556}
	r, _ := max(input...)
	fmt.Println(r)
	r, _ = min(input...)
	fmt.Println(r)
}

func max(n ...int) (int, error) {
	// check if there is at elast one element

	if len(n) == 0 {
		return 0, fmt.Errorf("[!]Not enough elment(s) provided\n")
	}

	// sort the slice first
	sort.Ints(n)
	return n[len(n)-1], nil
}

func min(n ...int) (int, error) {
	// check if there is at elast one element

	if len(n) == 0 {
		return 0, fmt.Errorf("[!]Not enough elment(s) provided\n")
	}

	// sort the slice first
	sort.Ints(n)
	return n[0], nil
}

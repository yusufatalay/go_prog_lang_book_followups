// rempdump removes adjacent duplicates from a slice
package main

import "fmt"

func remove(s *[]int, i int) {

	copy((*s)[i:], (*s)[i+1:])
	(*s) = (*s)[:len((*s))-1]
}

func remdup(s *[]int) {

	for i := 1; i < len(*s)-1; i++ {

		if (*s)[i] == (*s)[i-1] || (*s)[i] == (*s)[i+1] {
			remove(s, i)
		}

	}

}

func main() {
	s := []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 1, 4, 5, 6}

	remdup(&s)

	fmt.Println(s)
}

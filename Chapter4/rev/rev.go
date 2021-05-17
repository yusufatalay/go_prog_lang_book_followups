// reverses the given array in-place
package main

import "fmt"

func rev(s *[]int) {

	for i, j := 0, len((*s))-1; i < j; i, j = i+1, j-1 {

		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]

	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rev(&arr)

	fmt.Println(arr)

}

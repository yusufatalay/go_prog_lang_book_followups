package main

import "sort"

func isPalindrome(s sort.Interface) bool {
	slen := s.Len() - 1
	for i := 0; i < slen/2; i++ {
		if !s.Less(i, slen-i) && !s.Less(slen-i, i) {
			return false
		}
	}
	return true
}

func main() {
}

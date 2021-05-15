package main

import (
	"fmt"
	"strings"
)

// check if two given strings are anagram of eachother or not

func checkAnagram(s1, s2 string) bool {

	lenS := len(s1)
	if lenS != len(s2) {
		return false
	}

	for i, _ := range s1 {
		if strings.Count(s1, string(s2[i])) != strings.Count(s2, string(s1[i])) {
			return false
		}
	}
	return true

}

func main() {

	fmt.Println(checkAnagram("anan", "nanb"))
}

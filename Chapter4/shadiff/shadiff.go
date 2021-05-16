package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func shadiff(str1, str2 string) int {

	diffCount := 0
	var checked []byte

	cyp1 := sha256.Sum256([]byte(str1))
	cyp1slc := cyp1[:]

	cyp2 := sha256.Sum256([]byte(str2))
	cyp2slc := cyp2[:]

	for _, v := range cyp1slc {

		toSearch := []byte{v}
		if !bytes.Contains(cyp2slc, toSearch) && !bytes.Contains(checked, toSearch) {
			diffCount += 1
			checked = append(checked, v)
		}
	}

	return diffCount
}

func main() {

	str1 := "first string"
	str2 := "second string"

	fmt.Printf("Number of bits differ from sha256 of st1: \"%s\" and str2: \"%s\" is -> %d\n", str1, str2, shadiff(str1, str2))
}

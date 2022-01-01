//	charcount computes counts of Unicode characters in a string
package charcount

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// charcount returns a map that contains amount of occurence of each unicode character
func charcount(s string) map[rune]int {
	counts := make(map[rune]int)
	// counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0
	// count of invalid UTF-8 characters
	counter := 0
	in := bufio.NewReader(strings.NewReader(s))
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error

		if counter == len(s) {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		counter++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	return counts
}

package main

import (
	"bytes"
	"fmt"
	"strings"
)

// IntSet is a set of small non-negative integers
// Its zero value represents empty set
type IntSet struct {
	words []uint64
}

// Has reports wether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String reutrns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	str := s.String()
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	starr := strings.Split(str, " ")
	return len(starr)
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {

	word, bit := x/64, uint64(x%64)
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var d IntSet
	for _, word := range s.words {
		d.words = append(d.words, word)
	}
	return &d
}
func main() {
	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)
	//	fmt.Println(y.String())

	x.UnionWith(&y)
	z := x.Copy()
	fmt.Println(x.String())
	fmt.Println(z.String())
	//	fmt.Println(x.String())
	//	fmt.Println(y.String())
	//	fmt.Println(x.Has(9), x.Has(123))

}

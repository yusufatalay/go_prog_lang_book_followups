package intset

import (
	"bytes"
	"fmt"
	"strings"
)

// SYSTEM_BIT is check if the system has 64 or 32 bit architecture
// SYSTEM_BIT == 64 if system is 64 bit, 0 if system is 32 bit
const SYSTEM_BIT int = 32 << (^uint(0) >> 63)

// IntSet is a set of small non-negative integers
// Its zero value represents empty set
type IntSet struct {
	words []uint
}

// Has reports wether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	word, bit := x/sysbit, uint(x%sysbit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	word, bit := x/sysbit, uint(x%sysbit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all the elements to the set
func (s *IntSet) AddAll(xs ...int) {
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	for _, x := range xs {
		word, bit := x/sysbit, uint(x%sysbit)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
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

// IntersectWith sets s to the intersect of s and t
func (s *IntSet) InterSectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
			s.words[len(s.words)-1-i] &= 0
		}
	}
}

// DifferenceWith sets s to the difference between s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// neat little math trick here : !(p => q) only true if
			// p is 1 and q is 0
			s.words[i] = ^(^s.words[i] | tword)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = (s.words[i] ^ tword)
		}
	}
}

// String reutrns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < sysbit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", sysbit*i+j)
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
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	word, bit := x/sysbit, uint(x%sysbit)
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

// Elems returns a slice that contains the elements of s
func (s *IntSet) Elems() []int {
	sysbit := SYSTEM_BIT
	if SYSTEM_BIT == 0 {
		sysbit = 32
	}
	var result []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < sysbit; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, sysbit*i+j)
			}
		}
	}
	return result
}

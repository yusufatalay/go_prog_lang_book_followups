package charcount

import (
	"reflect"
	"testing"
)

type test struct {
	input string
	want  map[rune]int
}

func TestCharCount(t *testing.T) {

	var tests = []struct {
		input string
		want  map[rune]int
	}{
		{"aaaa", map[rune]int{'a': 4}},
		{"aabb", map[rune]int{'a': 2, 'b': 2}},
		{"abcd", map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1}},
	}

	for _, test := range tests {
		if got := charcount(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("charcount(%s) = %+v\n", test.input, got)
		}
	}

}

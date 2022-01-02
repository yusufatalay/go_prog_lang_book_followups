package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		input string
		delim string
		want  []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a b c", " ", []string{"a", "b", "c"}},
		/* Kinda all the same */
	}

	for _, test := range tests {
		if got := strings.Split(test.input, test.delim); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Split(%q,%q), returned %v+, want %v+", test.input, test.delim, got, test.want)
		}
	}
}

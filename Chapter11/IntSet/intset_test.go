package intset

import (
	"testing"
)

func TestHas(t *testing.T) {

	var intset IntSet

	intset.AddAll(1, 2, 999, 3, 787, 22)

	var tests = []struct {
		input int
		want  bool
	}{
		{4, false},
		{1, true},
		{999, true},
	}

	for _, test := range tests {
		if got := intset.Has(test.input); test.want != got {
			t.Errorf("Has(%d) = %v\n", test.input, got)
		}
	}

}

func TestAdd(t *testing.T) {
	var intset IntSet

	var tests = []struct {
		input int
		want  bool
	}{
		{3, true},
		{0, true},
	}

	for _, test := range tests {
		intset.Add(test.input)
		if got := intset.Has(test.input); got != test.want {
			t.Errorf("List after Add(%d), %+v", test.input, intset)
		}
	}

}

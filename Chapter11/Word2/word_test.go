package word

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan , a canal: Panama")
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v ", test.input, got)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contends
// are derived from the pseudo-random generator rng
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 5 // random length up to 29 at least 6 characters
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i += 2 {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		// insert nonletter characters after evey letter in here
		runes[i+1] = ','
		runes[n-1-i] = r
	}
	return string(runes)
}

// randomNonPalindrome returns a non-palindrome whose length and contends
// are derived from the pseudo-random generator rng
func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 5 // random length up to 29 at least 6 characters
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
	}
	return string(runes)
}
func TestRandomNonPalindrome(t *testing.T) {
	// Initialize a pseudo-random numer generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

func TestRandomPalindrome(t *testing.T) {
	// Initialize a pseudo-random numer generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

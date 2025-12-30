package google_test

import (
	"gocoder/google"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
		{"1[a]", "a"},
		{"", ""},
		{"a", "a"},
		{"2[y]3[x]ef", "yyxxxef"},
	}

	for _, test := range tests {
		result := google.DecodeString(test.s)
		if result != test.expected {
			t.Errorf("For string %s, expected %s, but got %s", test.s, test.expected, result)
		}
	}
}

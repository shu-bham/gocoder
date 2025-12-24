package google

import "testing"

func TestLengthOfLongestSubstring_BasicExample1(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_AllSameChars(t *testing.T) {
	s := "bbbbb"
	expected := 1
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_TrickyExample(t *testing.T) {
	s := "pwwkew"
	expected := 3
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_EmptyString(t *testing.T) {
	s := ""
	expected := 0
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_TwoRepeatingChars(t *testing.T) {
	s := "aab"
	expected := 2
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_NoRepeatingChars(t *testing.T) {
	s := "dvdf"
	expected := 3
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_MixedRepeating(t *testing.T) {
	s := "abacadaea"
	expected := 3
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstring_SingleSpace(t *testing.T) {
	s := " "
	expected := 1
	result := LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

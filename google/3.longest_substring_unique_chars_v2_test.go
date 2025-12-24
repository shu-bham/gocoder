package google

import "testing"

func TestLengthOfLongestSubstringV2_BasicExample1(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_AllSameChars(t *testing.T) {
	s := "bbbbb"
	expected := 1
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_TrickyExample(t *testing.T) {
	s := "pwwkew"
	expected := 3
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_EmptyString(t *testing.T) {
	s := ""
	expected := 0
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_TwoRepeatingChars(t *testing.T) {
	s := "aab"
	expected := 2
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_NoRepeatingChars(t *testing.T) {
	s := "dvdf"
	expected := 3
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_MixedRepeating(t *testing.T) {
	s := "abacadaea"
	expected := 3
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_SingleSpace(t *testing.T) {
	s := " "
	expected := 1
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

func TestLengthOfLongestSubstringV2_(t *testing.T) {
	s := "abba"
	expected := 2
	result := LengthOfLongestSubstringV2(s)
	if result != expected {
		t.Errorf("For string '%s', expected %d, but got %d", s, expected, result)
	}
}

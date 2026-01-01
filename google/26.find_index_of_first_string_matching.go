package google

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}

	}
	return -1
}

// computeLPS builds the Longest Proper Prefix which is also Suffix array
// This tells us, for each position, how much of the start of the pattern
// matches the end of the substring up to that position
func computeLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)

	// First position is always 0 (no proper prefix exists)
	lps[0] = 0

	// len tracks the length of the current matching prefix-suffix
	// It's also the index in the pattern we're comparing against
	len := 0

	// i is our current position in the pattern we're analyzing
	i := 1

	for i < m {
		if pattern[i] == pattern[len] {
			// Characters match! We've extended our matching prefix-suffix
			// Example: pattern="ABAB", i=3, len=1
			// pattern[3]='B' matches pattern[1]='B'
			// So now we have "AB" matching at both ends
			len++
			lps[i] = len
			i++
		} else {
			// Mismatch! We need to "fall back"
			if len != 0 {
				// Don't give up completely - use the LPS value to fall back
				// This is the KEY INSIGHT: we can skip ahead because we know
				// some prefix is already matching from our previous work
				// Example: if len=2, we had "AB" matching
				// lps[len-1] tells us how much of that "AB" overlaps with the start
				len = lps[len-1]

				// Note: we DON'T increment i here
				// We'll retry the same position with our new, shorter matching length
			} else {
				// len == 0 means we're at the start, nothing matches
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// KMPSearch finds all occurrences of pattern in text
// Returns the starting indices of all matches
func KMPSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	if m == 0 || m > n {
		return []int{}
	}

	// Step 1: Build the LPS array for the pattern
	// This preprocesses the pattern to find self-overlaps
	lps := computeLPS(pattern)

	matches := []int{}

	// i tracks our position in the text
	// j tracks our position in the pattern
	i := 0
	j := 0

	for i < n {
		if pattern[j] == text[i] {
			// Characters match! Advance both pointers
			i++
			j++
		}

		if j == m {
			// We've matched the entire pattern!
			// Record the starting position (i - j gives us the start)
			matches = append(matches, i-j)

			// Can we find more matches? Use LPS to continue efficiently
			// Instead of starting over, we can skip ahead because we know
			// the last lps[j-1] characters of our match are the same as
			// the first lps[j-1] characters of the pattern
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			// Mismatch! This is where KMP shines
			if j != 0 {
				// We don't move i backwards! That's the magic of KMP
				// Instead, we use LPS to jump j to a position where we know
				// the prefix will match what we've already seen in the text
				//
				// Example: if we matched "ABAB" and then mismatched,
				// lps[3]=2 tells us the last "AB" is same as first "AB"
				// So we can continue checking from pattern[2]
				j = lps[j-1]
			} else {
				// j == 0 means even the first character didn't match
				// Just move to the next character in the text
				i++
			}
		}
	}

	return matches
}

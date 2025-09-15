package hard

import "strings"

/**
LC-68 https://leetcode.com/problems/text-justification

Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.

Company Tags:
'Airbnb', 'Amazon', 'Apple', 'Box', 'Coursera', 'Facebook', 'Google', 'Indeed', 'Intuit', 'Linkedin', 'Lyft', 'Microsoft', 'Paypal', 'Pinterest', 'Robinhood', 'Snapchat', 'Twilio', 'Twitter', 'Uber'
*/

func FullJustify(words []string, maxWidth int) []string {
	currLine, length := make([]string, 0), 0
	var res []string

	i := 0
	for i < len(words) {
		word := words[i]
		// if word cannot come in curr line then populate current line with spaces and append to result
		// len(currLine) -> no of single space
		if length+len(currLine)+len(word) > maxWidth {
			extraSpaces := maxWidth - length
			if len(currLine) == 1 {
				addSpace(&currLine[0], extraSpaces)
				res = append(res, currLine[0])
			} else {
				// for k words, we'll need to fill k-1 gaps between them
				spacePerWord := extraSpaces / (len(currLine) - 1)
				remainSpaces := extraSpaces % (len(currLine) - 1)

				for j := 0; j < len(currLine)-1; j++ {
					addSpace(&currLine[j], spacePerWord)
					if remainSpaces > 0 {
						addSpace(&currLine[j], 1)
						remainSpaces--
					}
				}

				join := strings.Join(currLine, "")
				res = append(res, join)
			}

			// reset currLine
			currLine, length = make([]string, 0), 0
		}

		// if word can come in curr line
		currLine = append(currLine, word)
		length += len(word)
		i++

	}

	// handle lastline
	lastLine := strings.Join(currLine, " ") // join with a single space
	trailingSpaces := maxWidth - len(lastLine)
	addSpace(&lastLine, trailingSpaces)
	res = append(res, lastLine)

	return res
}

func addSpace(s *string, times int) {
	for i := 0; i < times; i++ {
		*s += " "
	}
}

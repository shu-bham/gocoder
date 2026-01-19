package google

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	currLine, length := make([]string, 0), 0

	var res []string
	for _, word := range words {
		if length+len(currLine)+len(word) > maxWidth {
			// text justify curr line
			extraSpaces := maxWidth - length
			if len(currLine) == 1 {
				appendSpaces(&currLine[0], extraSpaces)
				res = append(res, currLine[0])
			} else {
				spaceCount := len(currLine) - 1
				spacePerWord := extraSpaces / spaceCount
				remSpaces := extraSpaces % spaceCount
				for j := 0; j < len(currLine)-1; j++ {
					appendSpaces(&currLine[j], spacePerWord)
					if remSpaces > 0 {
						appendSpaces(&currLine[j], 1)
						remSpaces--
					}
				}
				join := strings.Join(currLine, "")
				res = append(res, join)
			}
			currLine, length = []string{}, 0
		}

		// insert word in next line
		currLine = append(currLine, word)
		length += len(word)
	}

	lastLine := strings.Join(currLine, " ")
	extraSpaces := maxWidth - len(lastLine)
	appendSpaces(&lastLine, extraSpaces)
	res = append(res, lastLine)
	return res
}

func appendSpaces(s *string, spaces int) {
	for i := 0; i < spaces; i++ {
		*s += " "
	}
}

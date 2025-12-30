package google

import (
	"unicode"
)

func DecodeString(s string) string {
	stack := []interface{}{}
	currNum := 0
	currStr := ""

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			currStr += string(ch)
		} else if unicode.IsNumber(ch) {
			currNum = (currNum * 10) + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, currStr)
			stack = append(stack, currNum)
			currNum = 0
			currStr = ""
		} else if ch == ']' {
			num := stack[len(stack)-1].(int)
			stack = stack[:len(stack)-1]

			prevStr := stack[len(stack)-1].(string)
			stack = stack[:len(stack)-1]

			temp := ""
			for i := 0; i < num; i++ {
				temp += currStr
			}
			currStr = prevStr + temp
		}
	}
	return currStr
}

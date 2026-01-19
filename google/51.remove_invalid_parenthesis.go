package google

func removeInvalidParentheses(s string) []string {
	// store unmathced count of left and right parenthesis
	unmatchedLeft, unmatchedRight := 0, 0
	for _, ch := range s {
		if ch == '(' {
			unmatchedLeft++
		} else if ch == ')' {
			if unmatchedLeft > 0 {
				unmatchedLeft--
			} else {
				unmatchedRight++
			}
		}
	}
	var dfs func(str string, pos, left, right int)
	var result []string
	dfs = func(str string, pos, left, right int) {
		if left == 0 && right == 0 && isValid(str) {
			result = append(result, str)
			return
		}
		for i := pos; i < len(str); i++ {
			if i != pos && str[i] == str[i-1] {
				continue
			}
			newStr := str[:i] + str[i+1:]
			if left > 0 && str[i] == '(' {
				dfs(newStr, i, left-1, right)
			}
			if right > 0 && str[i] == ')' {
				dfs(newStr, i, left, right-1)
			}
		}
	}
	dfs(s, 0, unmatchedLeft, unmatchedRight)
	return result
}

func isValid(s string) bool {
	ctr := 0
	for _, ch := range s {
		if ch == '(' {
			ctr++
		} else if ch == ')' {
			ctr--
		}

		if ctr < 0 {
			return false
		}
	}
	if ctr > 0 {
		return false
	} else {
		return true
	}
}

func removeInvalidParenthesesTLE(s string) []string {
	if len(s) == 1 {
		return []string{""}
	}
	uniqueMap := make(map[string]bool)
	var ans []string
	var backtrack func(str string, count int)
	backtrack = func(str string, count int) {
		if count == 0 {
			if minParenthesisToRemoveToMakeValidString(str) == 0 {
				if uniqueMap[str] == false {
					uniqueMap[str] = true
					ans = append(ans, str)
				}
			}
		}
		for i := 0; i < len(str); i++ {
			if str[i] == '(' || str[i] == ')' {
				// remove ith char
				newStr := str[:i] + str[i+1:]
				backtrack(newStr, count-1)
			}
		}
	}
	count := minParenthesisToRemoveToMakeValidString(s)
	backtrack(s, count)
	return ans
}

func minParenthesisToRemoveToMakeValidString(s string) int {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			stackLen := len(stack)
			if stackLen == 0 {
				stack = append(stack, ch)
			} else {
				top := stack[stackLen-1]
				if top == '(' {
					stack = stack[:stackLen-1]
				} else {
					stack = append(stack, ch)
				}
			}
		}
	}
	return len(stack)
}

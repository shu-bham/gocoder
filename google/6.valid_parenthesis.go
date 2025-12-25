package google

func IsValid(s string) bool {
	stack := []int32{}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' {
			top := len(stack) - 1
			if top >= 0 && stack[top] == '(' {
				stack = stack[:top]
			} else {
				return false
			}
		} else if ch == ']' {
			top := len(stack) - 1
			if top >= 0 && stack[top] == '[' {
				stack = stack[:top]
			} else {
				return false
			}
		} else if ch == '}' {
			top := len(stack) - 1
			if top >= 0 && stack[top] == '{' {
				stack = stack[:top]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

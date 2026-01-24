package google

import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	j := 0

	for i := 0; i < n; {
		count := 0
		currCh := chars[i]
		for i < n && chars[i] == currCh {
			count++
			i++
		}

		if count == 1 {
			chars[j] = currCh
			j++
		} else {
			chars[j] = currCh
			j++
			for _, digit := range []byte(strconv.Itoa(count)) {
				chars[j] = digit
				j++
			}
		}
	}
	return j
}

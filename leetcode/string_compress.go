package leetcode

import (
	"fmt"
	"strconv"
)

// uses extra space
func Compress(chars []byte) int {
	prevCh := chars[0]
	ctr := 1
	var res []byte
	for i := 1; i < len(chars); i++ {
		currCh := chars[i]
		if currCh != prevCh {
			if ctr == 1 {
				res = append(res, prevCh)
			} else {
				res = append(res, prevCh)
				itoa := strconv.Itoa(ctr)
				res = append(res, itoa...)
			}
			prevCh = currCh
			ctr = 1
		} else {
			ctr++
		}
	}

	if ctr == 1 {
		res = append(res, prevCh)
	} else {
		res = append(res, prevCh)
		itoa := strconv.Itoa(ctr)
		res = append(res, itoa...)
	}

	for i, b := range res {
		chars[i] = b
	}

	//fmt.Println(string(res))

	return len(res)
}

// constant space solution
// inplace assignment
func CompressV2(chars []byte) int {
	pos := 0

	prevCh, ctr := chars[0], 1

	for i := 1; i < len(chars); i++ {
		currCh := chars[i]
		if currCh == prevCh {
			ctr++
		} else {
			// insert prevCh, ctr @ pos
			if ctr == 1 {
				chars[pos] = prevCh
				pos++
			} else {
				chars[pos] = prevCh
				pos++
				count := strconv.Itoa(ctr)
				for _, ch := range count {
					chars[pos] = byte(ch)
					pos++
				}
			}
			prevCh = currCh
			ctr = 1
		}
	}

	// insert prevCh, ctr @ pos
	if ctr == 1 {
		chars[pos] = prevCh
		pos++
	} else {
		chars[pos] = prevCh
		pos++
		count := strconv.Itoa(ctr)
		for _, ch := range count {
			chars[pos] = byte(ch)
			pos++
		}
	}

	fmt.Println(string(chars[:pos]))

	return pos
}

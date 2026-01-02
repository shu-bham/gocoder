package google

func HasSameDigits(s string) bool {
	bytes := []byte(s)
	for len(bytes) > 2 {
		for i := 1; i < len(bytes); i++ {
			bytes[i-1] = (bytes[i-1] + bytes[i]) % 10
		}
		bytes = bytes[:len(bytes)-1]
	}

	return bytes[0] == bytes[1]
}

// binomial coeffi
func HasSameDigitsv2(s string) bool {
	var nums []int
	for _, i := range s {
		nums = append(nums, int(i-'0'))
	}

	var binomialCoeff []int64

	n := len(nums)
	for i := 0; i < n-1; i++ {
		binomialCoeff = append(binomialCoeff, nCr(n-2, i))
	}

	var left, right int64
	for i := 0; i < n-1; i++ {
		left += int64(nums[i]) * binomialCoeff[i]
	}

	for i := 0; i < n-1; i++ {
		right += int64(nums[i+1]) * binomialCoeff[i]
	}

	return (left % 10) == (right % 10)

}

func nCr(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r // take advantage of symmetry
	}

	var result int64 = 1
	for i := 0; i < r; i++ {
		result = result * int64(n-i) / int64(i+1)
	}
	return result
}

package leetcode

// Medium: https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three
// analyse ternary form (base 3) of the no.
// if any coeffiecient is 2 then distinct powers can't make the number
func CheckPowersOfThree(n int) bool {
	for n > 0 {
		rem := n % 3
		if rem == 2 {
			return false
		}
		n = n / 3
	}

	return true
}

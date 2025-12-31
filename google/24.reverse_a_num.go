package google

import (
	"math"
)

func Reverse(x int) int {
	reversed := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		// Positive overflow check
		if reversed > math.MaxInt32/10 ||
			(reversed == math.MaxInt32/10 && digit > 7) {
			return 0
		}

		// Negative overflow check
		if reversed < math.MinInt32/10 ||
			(reversed == math.MinInt32/10 && digit < -8) {
			return 0
		}

		reversed = reversed*10 + digit
	}

	return reversed
}

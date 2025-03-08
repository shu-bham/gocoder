package leetcode

import (
	"fmt"
	"gocoder/util"
	"math"
)

func ClosestPrimes(left int, right int) []int {
	prevPrime := -1
	diff := math.MaxInt
	a, b := -1, -1
	for i := left; i <= right; i++ {
		if util.IsPrime(i) {
			if prevPrime == -1 {
				prevPrime = i
			} else {
				currDiff := i - prevPrime
				if currDiff < diff {
					a, b = prevPrime, i
					diff = currDiff
				}
				fmt.Println(prevPrime, i, a, b, currDiff, diff)
				// major optimisation:
				// once diff reach 2 no need to further look for prime pairs
				if currDiff == 2 {
					break
				}
				prevPrime = i
			}
		}
	}

	return []int{a, b}
}

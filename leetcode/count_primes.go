package leetcode

import "fmt"

// Medium: https://leetcode.com/problems/count-primes/
func CountPrimes(n int) int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
			fmt.Println(i, primes)
		}
	}

	ctr := 0
	for i := 2; i <= n; i++ {
		if primes[i] {
			ctr++
		}
	}
	return ctr
}

const N = 5_000_002

var primes []bool
var done bool
var count []int

func init() {
	primes = make([]bool, N)
	for i := 2; i < N; i++ {
		primes[i] = true
	}

	count = make([]int, N)
	for i := 2; i < N; i++ {
		count[i] = count[i-1]
		if !primes[i] {
			continue
		}
		for j := i * i; j < N; j += i {
			primes[j] = false
		}
		count[i]++
	}
}

func CountPrimesV2(n int) int {
	if n == 0 {
		return 0
	}
	return count[n-1]
}

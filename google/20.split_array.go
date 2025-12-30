package google

import "math"

// brute force appraoch - linear search
func SplitArray(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	lo := math.MinInt
	hi := 0
	for _, num := range nums {
		lo = max(lo, num) // minm sum for each subarray to receive atleast 1 element
		hi += num         // maxm sum when k == 1
	}

	for i := lo; i <= hi; i++ {
		p := findPartitions(nums, i)
		if p <= k {
			return i
		}
	}
	return -1
}

// optimised - binary search
func SplitArrayv2(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	lo := math.MinInt
	hi := 0
	for _, num := range nums {
		lo = max(lo, num) // minm sum for each subarray to receive atleast 1 element
		hi += num         // maxm sum when k == 1
	}

	for lo <= hi {
		mid := (lo + hi) / 2
		p := findPartitions(nums, mid)
		if p < k {
			hi = mid - 1
		} else if p > k {
			lo = mid + 1
		} else {
			// we are looking for minm
			hi = mid - 1
		}
	}
	return lo
}

func findPartitions(nums []int, sum int) int {
	partitionCtr := 1
	partitionSum := 0

	for _, num := range nums {
		if partitionSum+num <= sum {
			partitionSum += num
		} else {
			partitionCtr++
			partitionSum = num
		}
	}

	return partitionCtr
}

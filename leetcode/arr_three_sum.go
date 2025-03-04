package leetcode

import (
	"slices"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	n := len(nums)

	if n < 3 {
		return [][]int{}
	}

	var res [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < n-2; i++ {
		// Skip duplicate elements for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// Skip duplicate elements for j and k
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func ThreeSumV2(nums []int) [][]int {
	slices.Sort(nums)
	var n = len(nums)

	// This function starts at index i and moves right until it finds a new value.
	//	Purpose: Avoids duplicate elements when selecting the next number.
	var incr = func(i int) int {
		var val = nums[i]
		for i < n-2 && nums[i+1] == val {
			i += 1
		}
		return i + 1
	}

	var decr = func(i int) int {
		var val = nums[i]
		for i >= 2 && nums[i-1] == val {
			i -= 1
		}
		return i - 1
	}

	var results = make([][]int, 0)
	for i := 0; i < n-2; i = incr(i) {
		var a = nums[i]
		if a > 0 {
			return results
		}

		var j, k = i + 1, n - 1
		var b, c = nums[j], nums[k]
		for j < k {
			var sum = a + b + c

			if sum < 0 {
				j = incr(j)
				b = nums[j]
			} else if sum > 0 {
				k = decr(k)
				c = nums[k]
			} else {
				results = append(results, []int{a, b, c})
				j, k = incr(j), decr(k)
				b, c = nums[j], nums[k]
			}
		}
	}

	return results
}

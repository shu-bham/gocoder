package leetcode

import "fmt"

// Medium: https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i

// optimised: sliding window
func MinOperationsBinaryArrayV2(nums []int) int {
	ctr := 0
	n := len(nums) - 2
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			ctr++
			nums[i] = 1
			nums[i+1] = nums[i+1] ^ 1
			nums[i+2] = nums[i+2] ^ 1
		}
	}
	if nums[n]+nums[n+1] != 2 {
		return -1
	}

	return ctr
}

func MinOperationsBinaryArray(nums []int) int {
	left, right := 0, len(nums)-1

	minOps, ok := calcMinOps(nums, left, right)
	if ok {
		return minOps
	}

	return -1
}

func calcMinOps(nums []int, left int, right int) (int, bool) {
	fmt.Println(nums, left, right)
	if left > right {
		return 0, false
	}

	if nums[left] == 1 {
		if left == right {
			return 0, true
		}
		return calcMinOps(nums, left+1, right)
	} else {
		if right >= left+2 {
			newInd := -1
			flipBits(nums, left, left+2)
			for i := left; i <= right; i++ {
				if nums[i] == 0 {
					newInd = i
					break
				}
			}
			if newInd == -1 {
				return 1, true
			}
			ops, ok := calcMinOps(nums, newInd, right)
			if !ok {
				return 0, false
			}
			return 1 + ops, true
		} else {
			return 0, false
		}
	}
}

func flipBits(nums []int, left int, right int) {
	for i := left; i <= right; i++ {
		nums[i] = nums[i] ^ 1
	}
}

package leetcode

func MaximumCountPositiveOrNegative(nums []int) int {
	n := len(nums)

	// same sign
	if nums[0]*nums[n-1] > 0 {
		return n
	}

	// find first positive
	left, right := 0, n-1
	var firstPositive int = -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > 0 {
			firstPositive = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	//println(firstPositive, nums[firstPositive], n-firstPositive)

	// find first negative
	left, right = 0, n-1
	var firstnegative int = -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < 0 {
			firstnegative = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//println(firstnegative, nums[firstnegative], 1+firstnegative)

	if firstPositive == -1 && firstnegative == -1 {
		return 0
	} else if firstPositive == -1 {
		return 1 + firstnegative
	} else if firstnegative == -1 {
		return n - firstPositive
	} else {
		x := n - firstPositive
		y := firstnegative + 1
		return max(x, y)
	}

}

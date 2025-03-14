package leetcode

import "sort"

// Medium: https://leetcode.com/problems/maximum-candies-allocated-to-k-children
func MaximumCandies(candies []int, k int64) int {
	sum := int64(0)
	maxCandiesPerChild := int64(0)
	for _, c := range candies {
		sum += int64(c)
	}
	maxCandiesPerChild = sum / k

	sort.Slice(candies, func(i, j int) bool {
		return candies[i] > candies[j]
	})

	//left, right := int64(0), maxCandiesPerChild
	//
	//for left <= right {
	//	mid := (left + right) / 2
	//	ans := int64(0)
	//	for _, candy := range candies {
	//		ans += int64(candy) / mid
	//		if ans >= k {
	//			left = mid + 1
	//		} else {
	//			right = mid-1
	//		}
	//	}
	//}

	for candiesPerChild := maxCandiesPerChild; candiesPerChild >= 2; candiesPerChild-- {
		ans := int64(0)
		for _, c := range candies {
			ans += int64(c) / candiesPerChild
			if ans >= k {
				return int(candiesPerChild)
			}
		}
	}

	if sum >= k {
		return 1
	}

	return 0

}

func MaximumCandiesV2(candies []int, k int64) int {
	sum := 0
	maxCandiesPerChild := 0
	for _, c := range candies {
		sum += c
	}
	maxCandiesPerChild = sum / int(k)

	left, right := 0, maxCandiesPerChild

	for left < right {
		mid := (left + right + 1) / 2
		if candiesCanBeDistributed(candies, k, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return int(left)
}

func candiesCanBeDistributed(candies []int, k int64, mid int) bool {
	ans := 0
	for _, candy := range candies {
		ans += (candy) / mid
	}

	return ans >= int(k)
}

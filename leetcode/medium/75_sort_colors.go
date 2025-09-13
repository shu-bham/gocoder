package medium

/*
*
75. Sort Colors

https://leetcode.com/problems/sort-colors

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/
func SortColors(nums []int) {

	lo := 0 // insert `0` at this pos
	i := 0
	hi := len(nums) - 1 // insert `2` at this pos

	for i <= hi {
		if nums[i] == 0 {
			// swap nums[lo], nums[i]
			nums[lo], nums[i] = nums[i], nums[lo]
			i++
			lo++
		} else if nums[i] == 1 {
			i++
		} else {
			// swap nums[i], nums[hi]
			nums[i], nums[hi] = nums[hi], nums[i]
			hi--
		}
	}

}

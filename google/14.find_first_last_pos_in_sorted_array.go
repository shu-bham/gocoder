package google

func SearchRange(nums []int, target int) []int {
	none := []int{-1, -1}
	if len(nums) == 0 {
		return none
	}

	f := findFirst(nums, target)
	l := findLast(nums, target)

	if f == -1 && l == -1 {
		return none
	}

	return []int{f, l}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

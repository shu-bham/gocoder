package leetcode

// Medium: https://leetcode.com/problems/house-robber-iv
func MinCapability(nums []int, k int) int {
	maxReward := 0
	minReward := 1
	for _, num := range nums {
		maxReward = max(maxReward, num)
	}

	for minReward < maxReward {
		midReward := (minReward + maxReward) / 2
		possibleThefts := 0

		for i := 0; i < len(nums); i++ {
			if nums[i] <= midReward {
				possibleThefts++
				i++
			}
		}

		if possibleThefts >= k {
			maxReward = midReward
		} else {
			minReward = midReward + 1
		}
	}

	return minReward
}

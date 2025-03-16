package leetcode

import (
	"math"
	"slices"
)

// Medium: https://leetcode.com/problems/minimum-time-to-repair-cars
func RepairCars(ranks []int, cars int) int64 {
	mx := slices.Max(ranks)
	minTime, maxTime := int64(0), int64(mx)*int64(cars)*int64(cars)

	canAllCarsBeRepairedInGivenTime := func(v int64) bool {
		sum := int64(0)

		for _, rank := range ranks {
			sum += int64(math.Sqrt(float64(v / int64(rank))))
			if sum >= int64(cars) {
				return true
			}
		}

		return false
	}

	for minTime <= maxTime {
		midTime := (minTime + maxTime) / 2
		if canAllCarsBeRepairedInGivenTime(midTime) {
			maxTime = midTime - 1
		} else {
			minTime = midTime + 1
		}
	}

	return minTime

}

package google

import (
	"fmt"
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	result := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			result[k] = nums1[i]
			i++
			k++
		} else {
			result[k] = nums2[j]
			j++
			k++
		}
	}

	for i < m {
		result[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		result[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(result)

	if (m+n)%2 == 1 {
		return float64(result[(m+n)/2])
	} else {
		return (float64(result[(m+n-1)/2]) + float64(result[(m+n)/2])) / 2
	}

}

func FindMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0

	var val1, val2 int
	for ctr := 0; ctr <= (m+n)/2; ctr++ {
		val2 = val1
		if i != m && j != n {
			if nums1[i] < nums2[j] {
				val1 = nums1[i]
				i++
			} else {
				val1 = nums2[j]
				j++
			}
		} else if i < m {
			val1 = nums1[i]
			i++
		} else {
			val1 = nums2[j]
			j++
		}
	}

	println(val1, val2)

	if (m+n)%2 == 1 {
		return float64(val1)
	} else {
		return float64(val1+val2) / 2
	}

}

func FindMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return FindMedianSortedArraysV3(nums2, nums1)
	}

	n := n1 + n2
	leftHalfSize := (n + 1) / 2
	lo, hi := 0, n1

	for lo <= hi {
		mid1 := (lo + hi) / 2
		mid2 := leftHalfSize - mid1

		l1 := math.MinInt32
		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}
		r1 := math.MaxInt32
		if mid1 < n1 {
			r1 = nums1[mid1]
		}

		l2 := math.MinInt32
		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}
		r2 := math.MaxInt32
		if mid2 < n2 {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return float64(max(l1, l2))
			}
			return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2.0
		} else if l1 > r2 {
			hi = mid1 - 1
		} else {
			lo = mid1 + 1
		}
	}
	return 0
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

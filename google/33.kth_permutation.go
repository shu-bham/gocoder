package google

import (
	"slices"
	"strconv"
	"strings"
)

func GetPermutation(n int, k int) string {
	var arr []int

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < k-1; i++ {
		arr = nextPermutation(arr)
	}

	ans := ""
	for _, num := range arr {
		ans += strconv.Itoa(num)
	}
	return ans

	//allPermutations := permute(arr)
	//ans := ""
	//for _, num := range allPermutations[k-1] {
	//	ans += strconv.Itoa(num)
	//}
	//return ans

}

// optimised
func GetPermutationOpt(n int, k int) string {
	var arr []int
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
		fact[i] = i * fact[i-1]
	}
	k--
	sb := strings.Builder{}
	for i := n; i > 0; i-- {
		index := k / fact[i-1]
		k = k % fact[i-1]
		sb.WriteString(strconv.Itoa(arr[index]))
		arr = append(arr[:index], arr[index+1:]...)
	}

	return sb.String()

}

func permute(arr []int) [][]int {
	var result [][]int
	used := make([]bool, len(arr))
	current := make([]int, 0, len(arr))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(arr) {
			perm := make([]int, len(arr))
			copy(perm, current)
			result = append(result, perm)
			return
		}
		for i, num := range arr {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, num)
			backtrack()
			used[i] = false
			current = current[:len(current)-1]
		}

	}
	backtrack()
	return result
}

func nextPermutation(arr []int) []int {
	leftPos := -1
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j-1] < arr[j] {
			leftPos = j - 1
			break
		}
	}

	if leftPos == -1 {
		slices.Reverse(arr)
		return arr
	}

	for j := len(arr) - 1; j > leftPos; j-- {
		if arr[j] > arr[leftPos] {
			arr[j], arr[leftPos] = arr[leftPos], arr[j]
			break
		}
	}
	slices.Reverse(arr[leftPos+1:])
	return arr
}

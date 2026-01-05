package google

import "math"

func RangeQuery(arr []int, queries [][]int) []int {
	segTree := make([]int, 4*len(arr))
	buildTree(arr, 0, 0, len(arr)-1, &segTree)
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = query(0, 0, len(arr)-1, q[0], q[1], segTree)
	}
	return res
}

func query(index int, lo int, hi int, l int, r int, segTree []int) int {
	if lo >= l && hi <= r {
		return segTree[index]
	}

	if lo > r || hi < l {
		return math.MaxInt
	}

	mid := (lo + hi) / 2
	leftRes := query(2*index+1, lo, mid, l, r, segTree)
	rightRes := query(2*index+2, mid+1, hi, l, r, segTree)
	return min(leftRes, rightRes)
}

func buildTree(arr []int, index int, left int, right int, segTree *[]int) {
	if left == right {
		(*segTree)[index] = arr[left]
		return
	}
	mid := (left + right) / 2
	buildTree(arr, 2*index+1, left, mid, segTree)
	buildTree(arr, 2*index+2, mid+1, right, segTree)
	(*segTree)[index] = min((*segTree)[2*index+1], (*segTree)[2*index+2])
}

package google

type SegmentTreeNode struct {
	left      *SegmentTreeNode
	right     *SegmentTreeNode
	minVal    int
	minimaPos []int
}

type SegmentTree struct {
	n    int
	root *SegmentTreeNode
}

func NewTreeNode(left *SegmentTreeNode, right *SegmentTreeNode, minVal int, pos []int) *SegmentTreeNode {
	if pos == nil {
		pos = []int{}
	}
	return &SegmentTreeNode{
		left:      left,
		right:     right,
		minVal:    minVal,
		minimaPos: pos,
	}
}

func NewSegmentTree(arr []int) *SegmentTree {
	st := &SegmentTree{
		n: len(arr),
	}
	st.root = st.build(arr, 0, st.n-1)
	return st
}

func (st *SegmentTree) build(arr []int, left int, right int) *SegmentTreeNode {
	if left == right {
		return NewTreeNode(nil, nil, arr[left], []int{left})
	}
	mid := (left + right) / 2
	leftChild := st.build(arr, left, mid)
	rightChild := st.build(arr, mid+1, right)

	node := &SegmentTreeNode{
		left:  leftChild,
		right: rightChild,
	}

	if leftChild.minVal < rightChild.minVal {
		node.minVal = leftChild.minVal
		node.minimaPos = append([]int{}, leftChild.minimaPos...)
	} else if leftChild.minVal > rightChild.minVal {
		node.minVal = rightChild.minVal
		node.minimaPos = append([]int{}, rightChild.minimaPos...)
	} else {
		node.minVal = rightChild.minVal
		node.minimaPos = append(leftChild.minimaPos, rightChild.minimaPos...)
	}
	return node

}

func (st *SegmentTree) Query(ltarget, rtarget int) (int, []int) {
	return st.query(ltarget, rtarget, 0, st.n-1, st.root)
}

func (st *SegmentTree) query(ltarget int, rtarget int, l int, r int, node *SegmentTreeNode) (int, []int) {
	if ltarget == l && rtarget == r {
		return node.minVal, node.minimaPos
	}
	mid := (l + r) / 2
	if rtarget <= mid {
		return st.query(ltarget, rtarget, l, mid, node.left)
	} else if ltarget > mid {
		return st.query(ltarget, rtarget, mid+1, r, node.right)
	} else {
		lminVal, lPos := st.query(ltarget, mid, l, mid, node.left)
		rminVal, rPos := st.query(mid+1, rtarget, mid+1, r, node.right)

		if lminVal < rminVal {
			return lminVal, lPos
		} else if rminVal < lminVal {
			return rminVal, rPos
		} else {
			return lminVal, append(lPos, rPos...)
		}
	}
}

type interval struct {
	left  int
	right int
	val   int
}

func MinNumberOperations(target []int) int {
	stree := NewSegmentTree(target)
	intervals := []interval{{0, len(target) - 1, 0}}
	operations := 0
	for len(intervals) > 0 {
		interval_ := intervals[len(intervals)-1]
		intervals = intervals[:len(intervals)-1]

		left, right, val := interval_.left, interval_.right, interval_.val
		if right < left {
			continue
		}
		intervalMin, pos := stree.Query(left, right)
		operations += (intervalMin - val)
		val = intervalMin
		for _, breakPoint := range pos {
			if breakPoint-1 >= left {
				intervals = append(intervals, interval{left, breakPoint - 1, val})
			}
			left = breakPoint + 1
		}
		if right >= left {
			intervals = append(intervals, interval{left, right, val})
		}
	}
	return operations
}

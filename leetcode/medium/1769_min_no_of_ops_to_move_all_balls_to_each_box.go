package medium

/**
1769. Minimum Number of Operations to Move All Balls to Each Box
https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.

In one operation, you can move one ball from a box to an adjacent box. Box i is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may be more than one ball in some boxes.

Return an array answer of size n, where answer[i] is the minimum number of operations needed to move all the balls to the ith box.

Each answer[i] is calculated considering the initial state of the boxes.
*/

func MinOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		var temp int
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			if boxes[j] == '1'-0 {
				ctr := j - i
				if ctr > 0 {
					temp += ctr
				} else {
					temp += -ctr
				}
			}
		}
		res[i] = temp
	}

	return res
}

func MinOperationsV2(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)

	ballsToLeftCount := 0  // ball count to the left for current index
	ballsToRightCount := 0 // // ball count to the right for current index

	movesFromLeft := 0  // no. of moves reqd to move all ball on left to current index
	movesFromRight := 0 // no. of moves reqd to move all ball on right to current index

	for i := 0; i < n; i++ {
		res[i] += movesFromLeft
		ballsToLeftCount += int(boxes[i] - '0')
		// for next iter
		movesFromLeft += ballsToLeftCount

		j := n - 1 - i
		res[j] += movesFromRight
		ballsToRightCount += int(boxes[j] - '0')
		// for next iter
		movesFromRight += ballsToRightCount
	}

	return res
}

package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		original := createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
		cloned := leetcode.CloneGraph(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t2", func(t *testing.T) {
		original := createGraph([][]int{})
		cloned := leetcode.CloneGraph(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t3", func(t *testing.T) {
		original := createGraph([][]int{{}})
		cloned := leetcode.CloneGraph(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t4", func(t *testing.T) {
		original := createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
		cloned := leetcode.CloneGraph(original)

		assert.True(t, isGraphEqual(original, cloned))
	})
}

func TestCloneGraphV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		original := createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
		cloned := leetcode.CloneGraphV2(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t2", func(t *testing.T) {
		original := createGraph([][]int{})
		cloned := leetcode.CloneGraphV2(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t3", func(t *testing.T) {
		original := createGraph([][]int{{}})
		cloned := leetcode.CloneGraphV2(original)

		assert.True(t, isGraphEqual(original, cloned))
	})

	t.Run("t4", func(t *testing.T) {
		original := createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
		cloned := leetcode.CloneGraphV2(original)

		assert.True(t, isGraphEqual(original, cloned))
	})
}

func isGraphEqual(node1, node2 *leetcode.CloneGraphNode) bool {
	visited := make(map[*leetcode.CloneGraphNode]*leetcode.CloneGraphNode)
	return dfsCompare(node1, node2, visited)
}

func dfsCompare(n1, n2 *leetcode.CloneGraphNode, visited map[*leetcode.CloneGraphNode]*leetcode.CloneGraphNode) bool {
	if n1 == nil || n2 == nil {
		return n1 == n2
	}

	if n1.Val != n2.Val {
		return false
	}

	if len(n1.Neighbors) != len(n2.Neighbors) {
		return false
	}

	visited[n1] = n2

	for i := range n1.Neighbors {
		if visited[n1.Neighbors[i]] == nil {
			if !dfsCompare(n1.Neighbors[i], n2.Neighbors[i], visited) {
				return false
			}
		} else if visited[n1.Neighbors[i]] != n2.Neighbors[i] {
			return false
		}
	}

	return true
}

func createGraph(adj [][]int) *leetcode.CloneGraphNode {
	if len(adj) == 0 {
		return nil
	}

	nodes := make(map[int]*leetcode.CloneGraphNode)
	for i := range adj {
		nodes[i+1] = &leetcode.CloneGraphNode{Val: i + 1, Neighbors: make([]*leetcode.CloneGraphNode, 0)}
	}

	for i, neighbours := range adj {
		for _, neighbour := range neighbours {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[neighbour])
		}
	}

	return nodes[1]
}

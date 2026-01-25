package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructMaximumBinaryTreeOptimised(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected *TreeNode1
	}{
		{
			name:  "Example 1",
			input: []int{3, 2, 1, 6, 0, 5},
			expected: &TreeNode1{
				Val: 6,
				Left: &TreeNode1{
					Val: 3,
					Right: &TreeNode1{
						Val: 2,
						Right: &TreeNode1{
							Val: 1,
						},
					},
				},
				Right: &TreeNode1{
					Val: 5,
					Left: &TreeNode1{
						Val: 0,
					},
				},
			},
		},
		{
			name:  "Empty array",
			input: []int{},
			expected: nil,
		},
		{
			name:  "Single element array",
			input: []int{5},
			expected: &TreeNode1{
				Val: 5,
			},
		},
		{
			name:  "Sorted increasing array",
			input: []int{1, 2, 3, 4, 5},
			expected: &TreeNode1{
				Val: 5,
				Left: &TreeNode1{
					Val: 4,
					Left: &TreeNode1{
						Val: 3,
						Left: &TreeNode1{
							Val: 2,
							Left: &TreeNode1{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name:  "Sorted decreasing array",
			input: []int{5, 4, 3, 2, 1},
			expected: &TreeNode1{
				Val: 5,
				Right: &TreeNode1{
					Val: 4,
					Right: &TreeNode1{
						Val: 3,
						Right: &TreeNode1{
							Val: 2,
							Right: &TreeNode1{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constructMaximumBinaryTreeOptimised(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

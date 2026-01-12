package google_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/google"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "LeetCode example 1",
			grid: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
		{
			name: "LeetCode example 2",
			grid: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			name: "Single element '0'",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "Single element '1'",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "All '1's 2x2",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 4,
		},
		{
			name: "All '0's",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "3x3 with 2x2 square",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '0'},
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := google.MaximalSquare(tc.grid)
			assert.Equal(t, tc.want, got)
		})
	}
}

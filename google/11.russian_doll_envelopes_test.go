package google

import (
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		name      string
		envelopes [][]int
		want      int
	}{
		{
			name: "Example 1",
			envelopes: [][]int{
				{5, 4}, {6, 4}, {6, 7}, {2, 3},
			},
			want: 3,
		},
		{
			name: "Example 2",
			envelopes: [][]int{
				{1, 1}, {1, 1}, {1, 1},
			},
			want: 1,
		},
		{
			name: "Example 3",
			envelopes: [][]int{
				{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1},
			},
			want: 4,
		},
		{
			name: "Example 4",
			envelopes: [][]int{
				{4, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxEnvelopes(tt.envelopes); got != tt.want {
				t.Errorf("MaxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxEnvelopesv2(tt.envelopes); got != tt.want {
				t.Errorf("MaxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

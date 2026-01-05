package google

import "testing"

func TestMinNumberOperations(t *testing.T) {
	tests := []struct {
		name   string
		target []int
		want   int
	}{
		{
			name:   "Example 1",
			target: []int{1, 2, 3},
			want:   3,
		},
		{
			name:   "Example 2",
			target: []int{3, 1, 2},
			want:   4,
		},
		{
			name:   "Example 3",
			target: []int{1, 2, 3, 2, 1},
			want:   3,
		},
		{
			name:   "Example 4",
			target: []int{3, 1, 5, 4, 2},
			want:   7,
		},
		{
			name:   "Single element",
			target: []int{5},
			want:   5,
		},
		//{
		//	name:   "Empty array",
		//	target: []int{},
		//	want:   0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinNumberOperations(tt.target); got != tt.want {
				t.Errorf("MinNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

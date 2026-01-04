package google

import "testing"

func TestRangeQuery(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{
			name:    "Basic Functionality",
			arr:     []int{1, 3, 2, 7, 9, 11},
			queries: [][]int{{0, 5}, {1, 3}, {2, 2}},
			want:    []int{1, 2, 2},
		},
		{
			name:    "Larger Array and Multiple Queries",
			arr:     []int{5, 2, 8, 1, 9, 3, 6, 4},
			queries: [][]int{{0, 7}, {0, 0}, {7, 7}, {1, 6}, {3, 5}},
			want:    []int{1, 5, 4, 1, 1},
		},
		{
			name:    "Array with negative numbers",
			arr:     []int{-5, -2, -8, -1, -9, -3, -6, -4},
			queries: [][]int{{0, 7}, {0, 0}, {7, 7}, {1, 6}, {3, 5}},
			want:    []int{-9, -5, -4, -9, -9},
		},
		{
			name:    "Array with all same numbers",
			arr:     []int{7, 7, 7, 7, 7},
			queries: [][]int{{0, 4}, {1, 3}, {2, 2}},
			want:    []int{7, 7, 7},
		},
		{
			name:    "Query on a single element array",
			arr:     []int{10},
			queries: [][]int{{0, 0}},
			want:    []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeQuery(tt.arr, tt.queries)
			if len(got) != len(tt.want) {
				t.Errorf("RangeQuery() got = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("RangeQuery() at index %d = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

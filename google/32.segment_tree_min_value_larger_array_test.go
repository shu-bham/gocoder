package google

import "testing"

func TestRangeQuery_LargerArrayMultipleQueries(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3, 6, 4}
	queries := [][]int{{0, 7}, {0, 0}, {7, 7}, {1, 6}, {3, 5}}
	want := []int{1, 5, 4, 1, 1}

	got := RangeQuery(arr, queries)
	if len(got) != len(want) {
		t.Errorf("RangeQuery() got = %v, want %v", got, want)
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("RangeQuery() at index %d = %v, want %v", i, got[i], want[i])
		}
	}
}

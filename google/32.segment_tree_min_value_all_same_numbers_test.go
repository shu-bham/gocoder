package google

import "testing"

func TestRangeQuery_AllSameNumbers(t *testing.T) {
	arr := []int{7, 7, 7, 7, 7}
	queries := [][]int{{0, 4}, {1, 3}, {2, 2}}
	want := []int{7, 7, 7}

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

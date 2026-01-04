package google

import "testing"

func TestRangeQuery_BasicFunctionality(t *testing.T) {
	arr := []int{1, 3, 2, 7, 9, 11}
	queries := [][]int{{0, 5}, {1, 3}, {2, 2}}
	want := []int{1, 2, 2}

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

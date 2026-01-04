package google

import "testing"

func TestRangeQuery_SingleElementArray(t *testing.T) {
	arr := []int{10}
	queries := [][]int{{0, 0}}
	want := []int{10}

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

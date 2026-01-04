package google

import "testing"

func TestMinNumberOperations_EmptyArray(t *testing.T) {
	target := []int{}
	want := 0
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

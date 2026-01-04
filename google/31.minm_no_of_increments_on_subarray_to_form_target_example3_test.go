package google

import "testing"

func TestMinNumberOperations_Example3(t *testing.T) {
	target := []int{1, 2, 3, 2, 1}
	want := 3
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

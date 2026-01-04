package google

import "testing"

func TestMinNumberOperations_Example2(t *testing.T) {
	target := []int{3, 1, 2}
	want := 4
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

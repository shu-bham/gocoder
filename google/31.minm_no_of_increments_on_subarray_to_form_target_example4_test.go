package google

import "testing"

func TestMinNumberOperations_Example4(t *testing.T) {
	target := []int{3, 1, 5, 4, 2}
	want := 7
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

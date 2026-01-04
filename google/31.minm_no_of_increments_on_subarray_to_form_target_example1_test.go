package google

import "testing"

func TestMinNumberOperations_Example1(t *testing.T) {
	target := []int{1, 2, 3}
	want := 3
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

package google

import "testing"

func TestMinNumberOperations_SingleElement(t *testing.T) {
	target := []int{5}
	want := 5
	if got := MinNumberOperations(target); got != want {
		t.Errorf("MinNumberOperations() = %v, want %v", got, want)
	}
}

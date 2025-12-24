package google_test

import (
	"gocoder/google"
	"testing"
)

func TestTrap(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		expected := 6
		if got := google.Trap(height); got != expected {
			t.Errorf("Trap(%v) = %d; want %d", height, got, expected)
		}
	})

	t.Run("t2", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		expected := 9
		if got := google.Trap(height); got != expected {
			t.Errorf("Trap(%v) = %d; want %d", height, got, expected)
		}
	})
}

func TestTrapV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		expected := 6
		if got := google.TrapV2(height); got != expected {
			t.Errorf("Trap(%v) = %d; want %d", height, got, expected)
		}
	})

	t.Run("t2", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		expected := 9
		if got := google.TrapV2(height); got != expected {
			t.Errorf("Trap(%v) = %d; want %d", height, got, expected)
		}
	})
}

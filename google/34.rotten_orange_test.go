package google

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	// Test case 1: All oranges can be rotted in 4 minutes.
	t.Run("All oranges rot in 4 minutes", func(t *testing.T) {
		grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
		want := 4
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 2: Impossible to rot all oranges, isolated fresh orange.
	t.Run("Isolated fresh orange remains", func(t *testing.T) {
		grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
		want := -1
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 3: No fresh oranges initially, time should be 0.
	t.Run("No fresh oranges", func(t *testing.T) {
		grid := [][]int{{0, 2}}
		want := 0
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 4: Single fresh orange, no rotten oranges to start.
	t.Run("Single fresh orange, no initial rotten", func(t *testing.T) {
		grid := [][]int{{1}}
		want := -1
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 5: Single empty cell.
	t.Run("Single empty cell", func(t *testing.T) {
		grid := [][]int{{0}}
		want := 0
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 6: All rotten oranges initially.
	t.Run("All oranges initially rotten", func(t *testing.T) {
		grid := [][]int{{2, 2, 2, 2, 2}}
		want := 0
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})

	// Test case 7: All fresh oranges, no rotten initially.
	t.Run("All fresh oranges, no initial rotten (impossible)", func(t *testing.T) {
		grid := [][]int{{1, 1, 1, 1, 1}}
		want := -1
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		if got := OrangesRotting(gridCopy); got != want {
			t.Errorf("OrangesRotting() = %v, want %v", got, want)
		}
	})
}

package google

import (
	"fmt"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{3, 1, "123"},
		{3, 2, "132"},
		{3, 3, "213"},
		{3, 4, "231"},
		{3, 5, "312"},
		{3, 6, "321"},
		{4, 9, "2314"},
		{1, 1, "1"},
		{2, 1, "12"},
		{2, 2, "21"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d,k=%d", tt.n, tt.k), func(t *testing.T) {
			got := GetPermutation(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("GetPermutation() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestGetPermutationOpt(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{3, 1, "123"},
		{3, 2, "132"},
		{3, 3, "213"},
		{3, 4, "231"},
		{3, 5, "312"},
		{3, 6, "321"},
		{4, 9, "2314"},
		{1, 1, "1"},
		{2, 1, "12"},
		{2, 2, "21"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d,k=%d", tt.n, tt.k), func(t *testing.T) {
			got := GetPermutationOpt(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("GetPermutationOpt() = %s, want %s", got, tt.want)
			}
		})
	}
}
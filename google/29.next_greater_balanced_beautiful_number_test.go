package google

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextBeautifulNumber(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{n: 0, want: 1},
		{n: 1, want: 22},
		{n: 21, want: 22},
		{n: 22, want: 122},
		{n: 100, want: 122},
		{n: 121, want: 122},
		{n: 122, want: 212},
		{n: 200, want: 212},
		{n: 212, want: 221},
		{n: 220, want: 221},
		{n: 221, want: 333},
		{n: 300, want: 333},
		{n: 1000, want: 1333},
		{n: 1332, want: 1333},
		{n: 1333, want: 3133},
		{n: 3000, want: 3133},
		{n: 3132, want: 3133},
		{n: 3313, want: 3331},
		{n: 3330, want: 3331},
		{n: 3331, want: 4444},
		{n: 4000, want: 4444},
		{n: 6000, want: 14444},
		{n: 15000, want: 22333},
		{n: 100000, want: 122333},
		{n: 666665, want: 666666},
		{n: 1000000, want: 1224444},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			got := NextBeautifulNumber(tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}
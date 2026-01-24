package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Simple case I",
			input: "I",
			want:  1,
		},
		{
			name:  "Simple case V",
			input: "V",
			want:  5,
		},
		{
			name:  "Simple case X",
			input: "X",
			want:  10,
		},
		{
			name:  "Subtraction case IV",
			input: "IV",
			want:  4,
		},
		{
			name:  "Subtraction case IX",
			input: "IX",
			want:  9,
		},
		{
			name:  "Complex case LVIII",
			input: "LVIII",
			want:  58,
		},
		{
			name:  "Complex case MCMXCIV",
			input: "MCMXCIV",
			want:  1994,
		},
		{
			name:  "Empty string",
			input: "",
			want:  0,
		},
		{
			name:  "III",
			input: "III",
			want:  3,
		},
		{
			name:  "VC",
			input: "IM",
			want:  999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

package google_test

import (
	"gocoder/google"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "negative number",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "number ending with 0",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "zero",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "positive overflow",
			args: args{x: 1534236469},
			want: 0,
		},
		{
			name: "negative overflow",
			args: args{x: -2147483648},
			want: 0,
		},
		{
			name: "another positive overflow",
			args: args{x: 2147483647},
			want: 0,
		},
		{
			name: "no overflow positive",
			args: args{x: 1463847412},
			want: 2147483641,
		},
		{
			name: "no overflow negative",
			args: args{x: -1463847412},
			want: -2147483641,
		},
		{
			name: "single digit",
			args: args{x: 5},
			want: 5,
		},
		{
			name: "single digit negative",
			args: args{x: -5},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := google.Reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

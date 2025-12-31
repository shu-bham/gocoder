package google_test

import (
	"gocoder/google"
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 3},
			},
			want: []int{2, 3, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := google.NextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

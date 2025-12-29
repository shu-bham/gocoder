package google_test

import (
	"gocoder/google"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals []google.Interval
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				intervals: []google.Interval{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				intervals: []google.Interval{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				intervals: []google.Interval{},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				intervals: []google.Interval{
					{1, 5},
					{8, 9},
					{8, 9},
				},
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				intervals: []google.Interval{
					{1, 5},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				intervals: []google.Interval{
					{1, 2},
					{2, 3},
					{3, 4},
				},
			},
			want: 1,
		},
		{
			name: "case 7",
			args: args{
				intervals: []google.Interval{
					{1, 10},
					{2, 3},
					{4, 5},
					{6, 7},
				},
			},
			want: 2,
		},
		{
			name: "case 8",
			args: args{
				intervals: []google.Interval{
					{1, 5},
					{2, 6},
					{3, 7},
				},
			},
			want: 3,
		},
		{
			name: "case 9",
			args: args{
				intervals: []google.Interval{
					{9, 10},
					{4, 9},
					{4, 17},
				},
			},
			want: 2,
		},
		{
			name: "case 10",
			args: args{
				intervals: []google.Interval{
					{2, 11},
					{6, 16},
					{11, 16},
				},
			},
			want: 2,
		},
		{
			name: "case 11 - more intervals",
			args: args{
				intervals: []google.Interval{
					{1, 18},
					{18, 23},
					{15, 29},
					{4, 15},
					{2, 11},
					{5, 13},
					{6, 17},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := google.MinMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := google.MinMeetingRoomsv2(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRoomsv2() = %v, want %v", got, tt.want)
			}
		})
	}
}

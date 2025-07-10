package main

import "testing"

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		startTime []int
		endTime   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				eventTime: 5,
				startTime: []int{1, 3},
				endTime:   []int{2, 5},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				eventTime: 10,
				startTime: []int{0, 7, 9},
				endTime:   []int{1, 8, 10},
			},
			want: 7,
		},
		{
			name: "case 3",
			args: args{
				eventTime: 10,
				startTime: []int{0, 3, 7, 9},
				endTime:   []int{1, 4, 8, 10},
			},
			want: 6,
		},
		{
			name: "case 4",
			args: args{
				eventTime: 5,
				startTime: []int{0, 1, 2, 3, 4},
				endTime:   []int{1, 2, 3, 4, 5},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

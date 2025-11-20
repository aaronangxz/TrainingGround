package main

import "testing"

func Test_intersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionSizeTwo(tt.args.intervals); got != tt.want {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

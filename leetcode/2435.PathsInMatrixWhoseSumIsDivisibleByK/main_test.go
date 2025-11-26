package main

import "testing"

func Test_numberOfPaths(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
				k:    3,
			},
			want: 2,
		}, {
			name: "case 2",
			args: args{
				grid: [][]int{{0, 0}},
				k:    5,
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPaths(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("numberOfPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

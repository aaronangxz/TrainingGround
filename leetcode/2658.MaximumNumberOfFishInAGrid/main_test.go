package main

import "testing"

func Test_findMaxFish(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{{4}},
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				grid: [][]int{{4, 0}},
			},
			want: 4,
		},
		{
			name: "case 5",
			args: args{
				grid: [][]int{{6, 1, 10}},
			},
			want: 17,
		},
		{
			name: "case 6",
			args: args{
				grid: [][]int{{3, 10, 5, 8}},
			},
			want: 26,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxFish(tt.args.grid); got != tt.want {
				t.Errorf("findMaxFish() = %v, want %v", got, tt.want)
			}
		})
	}
}

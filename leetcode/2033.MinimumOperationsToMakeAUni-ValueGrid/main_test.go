package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		grid [][]int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{2, 4}, {6, 8}},
				x:    2,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{1, 5}, {2, 3}},
				x:    1,
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{{1, 2}, {3, 4}},
				x:    2,
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.grid, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

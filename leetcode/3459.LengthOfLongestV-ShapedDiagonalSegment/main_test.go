package main

import "testing"

func Test_lenOfVDiagonal(t *testing.T) {
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
				grid: [][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}},
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				grid: [][]int{{1}},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenOfVDiagonal(tt.args.grid); got != tt.want {
				t.Errorf("lenOfVDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

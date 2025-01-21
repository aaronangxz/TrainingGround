package main

import "testing"

func Test_gridGame(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{2, 5, 4}, {1, 5, 1}},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{3, 3, 1}, {8, 5, 2}},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}},
			},
			want: 7,
		},
		{
			name: "case 4",
			args: args{
				grid: [][]int{{10, 12, 14, 19, 19, 12, 19, 2, 17}, {20, 7, 17, 14, 3, 1, 1, 17, 12}},
			},
			want: 58,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridGame(tt.args.grid); got != tt.want {
				t.Errorf("gridGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

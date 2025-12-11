package main

import "testing"

func Test_countCoveredBuildings(t *testing.T) {
	type args struct {
		matrixSide int
		buildings  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrixSide: 3,
				buildings:  [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				matrixSide: 3,
				buildings:  [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				matrixSide: 5,
				buildings:  [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCoveredBuildings(tt.args.matrixSide, tt.args.buildings); got != tt.want {
				t.Errorf("countCoveredBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n:         2,
				present:   []int{1, 2},
				future:    []int{4, 3},
				hierarchy: [][]int{{1, 2}},
				budget:    3,
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				n:         2,
				present:   []int{3, 4},
				future:    []int{5, 8},
				hierarchy: [][]int{{1, 2}},
				budget:    4,
			},
			want: 4,
		}, {
			name: "case 3",
			args: args{
				n:         3,
				present:   []int{4, 6, 8},
				future:    []int{7, 9, 11},
				hierarchy: [][]int{{1, 2}, {1, 3}},
				budget:    10,
			},
			want: 10,
		}, {
			name: "case 4",
			args: args{
				n:         3,
				present:   []int{5, 2, 3},
				future:    []int{8, 5, 6},
				hierarchy: [][]int{{1, 2}, {2, 3}},
				budget:    7,
			},
			want: 12,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.n, tt.args.present, tt.args.future, tt.args.hierarchy, tt.args.budget); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

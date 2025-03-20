package main

import (
	"reflect"
	"testing"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		query [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				n:     5,
				edges: [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}},
				query: [][]int{{0, 3}, {3, 4}},
			},
			want: []int{1, -1},
		},
		{
			name: "case 2",
			args: args{
				n:     3,
				edges: [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}},
				query: [][]int{{1, 2}},
			},
			want: []int{0},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.n, tt.args.edges, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_maxPoints(t *testing.T) {
	type args struct {
		grid    [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				grid:    [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}},
				queries: []int{5, 6, 2},
			},
			want: []int{5, 8, 1},
		},
		{
			name: "case 2",
			args: args{
				grid:    [][]int{{5, 2, 1}, {1, 1, 2}},
				queries: []int{3},
			},
			want: []int{0},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.grid, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_maxTargetNodes(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
				edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
			},
			want: []int{8, 7, 7, 8, 8},
		},
		{
			name: "case 2",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
				edges2: [][]int{{0, 1}, {1, 2}, {2, 3}},
			},
			want: []int{3, 6, 6, 6, 6},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTargetNodes(tt.args.edges1, tt.args.edges2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTargetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

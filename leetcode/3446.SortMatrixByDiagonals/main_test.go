package main

import (
	"reflect"
	"testing"
)

func Test_sortMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}},
			},
			want: [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}},
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{0, 1}, {1, 2}},
			},
			want: [][]int{{2, 1}, {1, 0}},
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{{1}},
			},
			want: [][]int{{1}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

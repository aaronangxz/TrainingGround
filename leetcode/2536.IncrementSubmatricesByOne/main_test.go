package main

import (
	"reflect"
	"testing"
)

func Test_rangeAddQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				n:       3,
				queries: [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}},
			},
			want: [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}},
		},
		{
			name: "case 2",
			args: args{
				n:       2,
				queries: [][]int{{0, 0, 1, 1}},
			},
			want: [][]int{{1, 1}, {1, 1}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeAddQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeAddQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

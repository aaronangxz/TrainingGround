package main

import "testing"

func Test_magnificentSets(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n:     6,
				edges: [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				n:     3,
				edges: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnificentSets(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("magnificentSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

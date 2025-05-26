package main

import "testing"

func Test_largestPathValue(t *testing.T) {
	type args struct {
		colors string
		edges  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				colors: "abaca",
				edges:  [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				colors: "a",
				edges:  [][]int{{0, 0}},
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPathValue(tt.args.colors, tt.args.edges); got != tt.want {
				t.Errorf("largestPathValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

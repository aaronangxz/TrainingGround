package main

import "testing"

func Test_minZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums:    []int{2, 0, 2},
				queries: [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums:    []int{4, 3, 2, 1},
				queries: [][]int{{1, 3, 2}, {0, 2, 1}},
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("minZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

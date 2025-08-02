package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				A: []int{4, 2, 2, 2},
				B: []int{1, 4, 1, 2},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				A: []int{2, 3, 4, 1},
				B: []int{3, 2, 5, 1},
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

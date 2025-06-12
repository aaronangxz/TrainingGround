package main

import "testing"

func Test_maxAdjacentDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 4},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{-5, -10, -5},
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAdjacentDistance(tt.args.nums); got != tt.want {
				t.Errorf("maxAdjacentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

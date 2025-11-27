package main

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2},
				k:    1,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{-1, -2, -3, -4, -5},
				k:    4,
			},
			want: -10,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{-5, 1, 2, -3, 4},
				k:    2,
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

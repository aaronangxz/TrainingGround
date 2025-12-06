package main

import "testing"

func Test_countPartitions(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{9, 4, 1, 3, 7},
				k:    4,
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 3, 4},
				k:    0,
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPartitions(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

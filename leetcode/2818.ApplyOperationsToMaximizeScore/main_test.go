package main

import "testing"

func Test_maximumScore(t *testing.T) {
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
				nums: []int{8, 3, 9, 3, 8},
				k:    2,
			},
			want: 81,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{19, 12, 14, 6, 10, 18},
				k:    3,
			},
			want: 4788,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

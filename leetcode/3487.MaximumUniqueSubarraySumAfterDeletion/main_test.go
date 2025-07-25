package main

import "testing"

func Test_maxSum(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 1, 0, 1, 1},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, -1, -2, 1, 0, -1},
			},
			want: 3,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

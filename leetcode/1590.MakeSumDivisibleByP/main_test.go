package main

import "testing"

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 1, 4, 2},
				p:    6,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{6, 3, 5, 2},
				p:    9,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3},
				p:    3,
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

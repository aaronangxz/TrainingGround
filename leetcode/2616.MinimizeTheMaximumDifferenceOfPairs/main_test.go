package main

import "testing"

func Test_minimizeMax(t *testing.T) {
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
				nums: []int{10, 1, 2, 7, 1, 3},
				p:    2,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{4, 2, 1, 2},
				p:    1,
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeMax(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minimizeMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

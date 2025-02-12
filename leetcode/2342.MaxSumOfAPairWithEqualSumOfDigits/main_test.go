package main

import "testing"

func Test_maximumSum(t *testing.T) {
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
				nums: []int{18, 43, 36, 13, 7},
			},
			want: 54,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{10, 12, 19, 14},
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{368, 369, 307, 304, 384, 138, 90, 279, 35, 396, 114, 328, 251, 364, 300, 191, 438, 467, 183},
			},
			want: 835,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum(tt.args.nums); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

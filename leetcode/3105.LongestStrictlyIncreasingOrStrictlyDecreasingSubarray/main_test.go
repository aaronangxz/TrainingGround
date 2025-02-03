package main

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
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
				nums: []int{1, 4, 3, 3, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 3, 3, 3},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

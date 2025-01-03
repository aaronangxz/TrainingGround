package main

import "testing"

func Test_waysToSplitArray(t *testing.T) {
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
				nums: []int{10, 4, -8, 7},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, 3, 1, 0},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 0},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{-2, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

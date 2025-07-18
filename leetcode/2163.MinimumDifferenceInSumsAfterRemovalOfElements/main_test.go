package main

import "testing"

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: -1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{7, 9, 5, 8, 1, 3},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

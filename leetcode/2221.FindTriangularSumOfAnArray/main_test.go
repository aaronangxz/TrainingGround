package main

import "testing"

func Test_triangularSum(t *testing.T) {
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
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{5},
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularSum(tt.args.nums); got != tt.want {
				t.Errorf("triangularSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

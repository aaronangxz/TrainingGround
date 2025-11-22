package main

import "testing"

func Test_minimumOperations(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 6, 9},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

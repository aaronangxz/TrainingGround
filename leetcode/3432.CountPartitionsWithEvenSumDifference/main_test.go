package main

import "testing"

func Test_countPartitions(t *testing.T) {
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
				nums: []int{10, 10, 3, 7, 6},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 2},
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{2, 4, 6, 8},
			},
			want: 3,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPartitions(tt.args.nums); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

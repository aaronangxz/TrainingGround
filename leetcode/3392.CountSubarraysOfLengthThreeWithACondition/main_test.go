package main

import "testing"

func Test_countSubarrays(t *testing.T) {
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
				nums: []int{1, 2, 1, 4, 1},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

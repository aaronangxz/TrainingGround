package main

import "testing"

func Test_kLengthApart(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
				k:    2,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 0, 0, 1, 0, 1},
				k:    2,
			},
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLengthApart(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}

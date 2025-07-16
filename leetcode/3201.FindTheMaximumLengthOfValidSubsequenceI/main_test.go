package main

import "testing"

func Test_maximumLength(t *testing.T) {
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
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 1, 1, 2, 1, 2},
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 3},
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

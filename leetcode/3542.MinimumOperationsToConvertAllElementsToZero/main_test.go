package main

import "testing"

func Test_minOperations(t *testing.T) {
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
				nums: []int{0, 2},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 1, 2, 1},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2},
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

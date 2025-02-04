package main

import "testing"

func Test_maxAscendingSum(t *testing.T) {
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
				nums: []int{10, 20, 30, 5, 10, 50},
			},
			want: 65,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{10, 20, 30, 40, 50},
			},
			want: 150,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{12, 17, 15, 13, 10, 11, 12},
			},
			want: 33,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{100, 10, 1},
			},
			want: 100,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAscendingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAscendingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

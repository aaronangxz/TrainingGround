package main

import "testing"

func Test_specialTriplets(t *testing.T) {
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
				nums: []int{6, 3, 6},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1, 0, 0},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{8, 4, 2, 8, 4},
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialTriplets(tt.args.nums); got != tt.want {
				t.Errorf("specialTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		apple    []int
		capacity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				apple:    []int{1, 3, 2},
				capacity: []int{4, 3, 1, 5, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				apple:    []int{5, 5, 5},
				capacity: []int{2, 4, 2, 7},
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBoxes(tt.args.apple, tt.args.capacity); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

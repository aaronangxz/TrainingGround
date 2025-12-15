package main

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				prices: []int{3, 2, 1, 4},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				prices: []int{8, 6, 7, 7},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				prices: []int{1},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.args.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}

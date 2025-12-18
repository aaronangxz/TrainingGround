package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices   []int
		strategy []int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				prices:   []int{4, 2, 8},
				strategy: []int{-1, 0, 1},
				k:        2,
			},
			want: 10,
		},
		{
			name: "case 2",
			args: args{
				prices:   []int{5, 4, 3},
				strategy: []int{1, 1, 0},
				k:        2,
			},
			want: 9,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.strategy, tt.args.k); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

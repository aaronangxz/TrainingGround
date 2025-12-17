package main

import "testing"

func Test_maximumProfit(t *testing.T) {
	type args struct {
		prices []int
		K      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				prices: []int{1, 7, 9, 8, 2},
				K:      2,
			},
			want: 14,
		},
		{
			name: "case 2",
			args: args{
				prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9},
				K:      3,
			},
			want: 36,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProfit(tt.args.prices, tt.args.K); got != tt.want {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

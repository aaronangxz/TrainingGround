package main

import "testing"

func Test_countGoodStrings(t *testing.T) {
	type args struct {
		low  int
		high int
		zero int
		one  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				low:  3,
				high: 3,
				zero: 1,
				one:  1,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				low:  2,
				high: 3,
				zero: 1,
				one:  2,
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodStrings(tt.args.low, tt.args.high, tt.args.zero, tt.args.one); got != tt.want {
				t.Errorf("countGoodStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

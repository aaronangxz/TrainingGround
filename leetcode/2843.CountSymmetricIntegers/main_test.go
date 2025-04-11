package main

import "testing"

func Test_countSymmetricIntegers(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				low:  1,
				high: 100,
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				low:  1200,
				high: 1230,
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSymmetricIntegers(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countSymmetricIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

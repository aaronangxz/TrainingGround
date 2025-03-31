package main

import "testing"

func Test_putMarbles(t *testing.T) {
	type args struct {
		weights []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				weights: []int{1, 3, 5, 1},
				k:       2,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				weights: []int{1, 3},
				k:       2,
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putMarbles(tt.args.weights, tt.args.k); got != tt.want {
				t.Errorf("putMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}

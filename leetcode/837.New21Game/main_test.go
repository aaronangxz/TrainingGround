package main

import "testing"

func Test_new21Game(t *testing.T) {
	type args struct {
		n      int
		k      int
		maxPts int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				n:      10,
				k:      1,
				maxPts: 10,
			},
			want: 1.00000,
		},
		{
			name: "case 2",
			args: args{
				n:      6,
				k:      1,
				maxPts: 10,
			},
			want: 0.60000,
		},
		{
			name: "case 3",
			args: args{
				n:      21,
				k:      17,
				maxPts: 10,
			},
			want: 0.73278,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new21Game(tt.args.n, tt.args.k, tt.args.maxPts); got != tt.want {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}

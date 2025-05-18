package main

import "testing"

func Test_colorTheGrid(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				m: 1,
				n: 1,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				m: 1,
				n: 2,
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				m: 5,
				n: 5,
			},
			want: 580986,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorTheGrid(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("colorTheGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

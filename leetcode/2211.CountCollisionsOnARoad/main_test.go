package main

import "testing"

func Test_countCollisions(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				directions: "RLRSLL",
			},
			want: 5,
		}, {
			name: "case 2",
			args: args{
				directions: "LLRR",
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCollisions(tt.args.directions); got != tt.want {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}

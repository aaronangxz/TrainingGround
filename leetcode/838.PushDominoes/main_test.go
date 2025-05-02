package main

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				dominoes: "RR.L",
			},
			want: "RR.L",
		},
		{
			name: "case 2",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
			want: "LL.RR.LLRRLL..",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}

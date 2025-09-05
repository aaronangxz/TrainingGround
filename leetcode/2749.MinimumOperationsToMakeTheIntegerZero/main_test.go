package main

import "testing"

func Test_makeTheIntegerZero(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n1: 3,
				n2: -2,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				n1: 5,
				n2: 7,
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeTheIntegerZero(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("makeTheIntegerZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

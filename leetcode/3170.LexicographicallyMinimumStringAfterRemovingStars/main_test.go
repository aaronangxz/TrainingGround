package main

import "testing"

func Test_clearStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "aaba*",
			},
			want: "aab",
		},
		{
			name: "case 2",
			args: args{
				s: "abc",
			},
			want: "abc",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearStars(tt.args.s); got != tt.want {
				t.Errorf("clearStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

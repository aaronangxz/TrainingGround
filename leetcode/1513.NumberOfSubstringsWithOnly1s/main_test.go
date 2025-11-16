package main

import "testing"

func Test_numSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "0110111",
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				s: "101",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "111111",
			},
			want: 21,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSub(tt.args.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

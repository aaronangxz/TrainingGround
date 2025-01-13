package main

import "testing"

func Test_minimumLength(t *testing.T) {
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
				s: "abaacbcbb",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "aa",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "aaabcbcbb",
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLength(tt.args.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

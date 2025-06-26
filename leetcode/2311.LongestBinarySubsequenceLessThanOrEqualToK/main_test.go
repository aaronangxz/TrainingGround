package main

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "1001010",
				k: 5,
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "00101001",
				k: 1,
			},
			want: 6,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

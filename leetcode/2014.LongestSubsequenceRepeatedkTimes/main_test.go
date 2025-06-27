package main

import "testing"

func Test_longestSubsequenceRepeatedK(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "letsleetcode",
				k: 2,
			},
			want: "let",
		},
		{
			name: "case 2",
			args: args{
				s: "bb",
				k: 2,
			},
			want: "b",
		},
		{
			name: "case 3",
			args: args{
				s: "a",
				k: 2,
			},
			want: "",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequenceRepeatedK(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubsequenceRepeatedK() = %v, want %v", got, tt.want)
			}
		})
	}
}

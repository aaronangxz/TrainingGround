package main

import "testing"

func Test_countPrefixSuffixPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				words: []string{"a", "aba", "ababa", "aa"},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				words: []string{"pa", "papa", "ma", "mama"},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				words: []string{"abab", "ab"},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixSuffixPairs(tt.args.words); got != tt.want {
				t.Errorf("countPrefixSuffixPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

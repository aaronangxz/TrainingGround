package main

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				word: "aabcaba",
				k:    0,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				word: "dabdcbdcdcd",
				k:    2,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				word: "aaabaaa",
				k:    2,
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}

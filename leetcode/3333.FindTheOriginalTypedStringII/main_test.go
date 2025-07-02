package main

import "testing"

func Test_possibleStringCount(t *testing.T) {
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
				word: "aabbccdd",
				k:    7,
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				word: "aabbccdd",
				k:    8,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				word: "aaabbb",
				k:    3,
			},
			want: 8,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

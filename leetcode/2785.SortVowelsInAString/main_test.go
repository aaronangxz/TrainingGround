package main

import "testing"

func Test_sortVowels(t *testing.T) {
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
				s: "lEetcOde",
			},
			want: "lEOtcede",
		},
		{
			name: "case 2",
			args: args{
				s: "lYmpH",
			},
			want: "lYmpH",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				word: "abbcccc",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				word: "abcd",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				word: "aaaa",
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_answerString(t *testing.T) {
	type args struct {
		word string
		n    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				word: "dbca",
				n:    2,
			},
			want: "dbc",
		},
		{
			name: "case 2",
			args: args{
				word: "gggg",
				n:    4,
			},
			want: "g",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerString(tt.args.word, tt.args.n); got != tt.want {
				t.Errorf("answerString() = %v, want %v", got, tt.want)
			}
		})
	}
}

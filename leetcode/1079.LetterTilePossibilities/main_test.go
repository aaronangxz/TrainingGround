package main

import "testing"

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				tiles: "AAB",
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				tiles: "AAABBC",
			},
			want: 188,
		},
		{
			name: "case 3",
			args: args{
				tiles: "V",
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

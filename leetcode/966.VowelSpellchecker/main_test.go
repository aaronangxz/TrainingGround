package main

import (
	"reflect"
	"testing"
)

func TestSpellchecker(t *testing.T) {
	type args struct {
		wordlist []string
		queries  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				wordlist: []string{"KiTe", "kite", "hare", "Hare"},
				queries:  []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			},
			want: []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{
			name: "case 2",
			args: args{
				wordlist: []string{"yellow"},
				queries:  []string{"YellOw"},
			},
			want: []string{"yellow"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spellchecker(tt.args.wordlist, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

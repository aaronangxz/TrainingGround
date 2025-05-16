package main

import (
	"reflect"
	"testing"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				words:  []string{"bab", "dab", "cab"},
				groups: []int{1, 2, 2},
			},
			want: []string{"bab", "dab"},
		},
		{
			name: "case 2",
			args: args{
				words:  []string{"a", "b", "c", "d"},
				groups: []int{1, 2, 3, 4},
			},
			want: []string{"a", "b", "c", "d"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordsInLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWordsInLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

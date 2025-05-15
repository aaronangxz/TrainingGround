package main

import (
	"reflect"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
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
				words:  []string{"c"},
				groups: []int{0},
			},
			want: []string{"c"},
		},
		{
			name: "case 2",
			args: args{
				words:  []string{"d"},
				groups: []int{1},
			},
			want: []string{"d"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

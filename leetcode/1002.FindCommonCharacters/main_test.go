package main

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				words: []string{"bella", "label", "roller"},
			},
			want: []string{"e", "l", "l"},
		},
		{
			name: "case 2",
			args: args{
				words: []string{"cool", "lock", "cook"},
			},
			want: []string{"c", "o"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

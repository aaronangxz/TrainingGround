package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicateFolder(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				paths: [][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}},
			},
			want: [][]string{{"d"}, {"d", "a"}},
		},
		{
			name: "case 2",
			args: args{
				paths: [][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}},
			},
			want: [][]string{{"a"}, {"a", "b"}, {"c"}, {"c", "b"}},
		},
		{
			name: "case 3",
			args: args{
				paths: [][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}},
			},
			want: [][]string{{"a"}, {"a", "b"}, {"c"}, {"c", "d"}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicateFolder(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

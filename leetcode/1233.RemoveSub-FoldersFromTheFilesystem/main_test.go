package main

import (
	"reflect"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	type args struct {
		folder []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			},
			want: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name: "case 2",
			args: args{
				folder: []string{"/a", "/a/b/c", "/a/b/d"},
			},
			want: []string{"/a"},
		},
		{
			name: "case 3",
			args: args{
				folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			},
			want: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSubfolders(tt.args.folder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

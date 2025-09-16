package main

import (
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{6, 4, 3, 2, 7, 6, 2},
			},
			want: []int{12, 7, 6},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, 2, 1, 1, 3, 3, 3},
			},
			want: []int{2, 1, 1, 3},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceNonCoprimes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

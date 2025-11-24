package main

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: []bool{true, false, false},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: []bool{false, false, false},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}

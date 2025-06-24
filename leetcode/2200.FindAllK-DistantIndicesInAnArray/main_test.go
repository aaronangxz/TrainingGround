package main

import (
	"reflect"
	"testing"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 4, 9, 1, 3, 9, 5},
				key:  9,
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				key:  2,
				k:    2,
			},
			want: []int{0, 1, 2, 3, 4},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

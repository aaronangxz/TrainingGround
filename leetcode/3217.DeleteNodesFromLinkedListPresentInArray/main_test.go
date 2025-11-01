package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"reflect"
	"testing"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
		head *common.TreeNodeWithNext
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNodeWithNext
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3},
				head: common.NewSliceToLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: common.NewSliceToLinkedList([]int{4, 5}),
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1},
				head: common.NewSliceToLinkedList([]int{1, 2, 1, 2, 1, 2}),
			},
			want: common.NewSliceToLinkedList([]int{2, 2, 2}),
		},
		{
			name: "case 3",
			args: args{
				nums: []int{5},
				head: common.NewSliceToLinkedList([]int{1, 2, 3, 4}),
			},
			want: common.NewSliceToLinkedList([]int{1, 2, 3, 4}),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedList(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

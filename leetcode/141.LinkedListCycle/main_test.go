package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *common.TreeNodeWithNext
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				head: common.NewSliceToCyclicLinkedList([]int{3, 2, 0, -4}, 1),
			},
			want: true,
		}, {
			name: "case 2",
			args: args{
				head: common.NewSliceToCyclicLinkedList([]int{1, 2}, 0),
			},
			want: true,
		}, {
			name: "case 3",
			args: args{
				head: common.NewSliceToCyclicLinkedList([]int{1}, -1),
			},
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

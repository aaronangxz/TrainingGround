package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *common.TreeNodeWithNext
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				head: common.NewSliceToLinkedList([]int{1, 0, 1}),
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				head: common.NewSliceToLinkedList([]int{0}),
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

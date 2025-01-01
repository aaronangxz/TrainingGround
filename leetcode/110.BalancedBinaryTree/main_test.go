package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				root: common.NewSliceToTree([]int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: true,
		}, {
			name: "case 2",
			args: args{
				root: common.NewSliceToTree([]int{1, 2, 2, 3, 3, 0, 0, 4, 4}),
			},
			want: false,
		}, {
			name: "case 3",
			args: args{
				root: common.NewSliceToTree([]int{}),
			},
			want: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

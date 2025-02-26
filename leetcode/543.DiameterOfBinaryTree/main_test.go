package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: common.NewSliceToTree([]int{1, 2, 3, 4, 5}),
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				root: common.NewSliceToTree([]int{1, 2}),
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

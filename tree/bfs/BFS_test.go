package bfs

import (
	"github.com/aaronangxz/TrainingGround/common"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	root := common.NewTreeNode(1)
	root.Left = common.NewTreeNode(2)
	root.Right = common.NewTreeNode(3)
	root.Left.Left = common.NewTreeNode(4)
	root.Left.Right = common.NewTreeNode(5)
	root.Right.Left = common.NewTreeNode(6)
	root.Right.Right = common.NewTreeNode(7)

	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*common.TreeNode
	}{
		{
			"Test",
			args{root: root},
			[]*common.TreeNode{root, root.Left, root.Right, root.Left.Left, root.Left.Right, root.Right.Left, root.Right.Right},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

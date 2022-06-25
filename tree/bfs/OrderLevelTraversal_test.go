package bfs

import (
	"github.com/aaronangxz/TrainingGround/common"
	"reflect"
	"testing"
)

func TestOrderLevelTraversal(t *testing.T) {
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
		want [][]int
	}{
		{
			"Test",
			args{root: root},
			[][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderLevelTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderLevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "case 1",
			args: args{common.NewSliceToTree([]int{4, 2, 7, 1, 3, 6, 9})},
			want: common.NewSliceToTree([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

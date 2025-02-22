package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"reflect"
	"testing"
)

func Test_recoverFromPreorder(t *testing.T) {
	type args struct {
		traversal string
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "case 1",
			args: args{
				traversal: "1-2--3--4-5--6--7",
			},
			want: common.NewSliceToTree([]int{1, 2, 5, 3, 4, 6, 7}),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverFromPreorder(tt.args.traversal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

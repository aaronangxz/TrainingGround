package main

import (
	"reflect"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
				queries:       [][]int{{0, 1}, {1, 0}},
			},
			want: []bool{false, true},
		}, {
			name: "case 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{},
				queries:       [][]int{{1, 0}, {0, 1}},
			},
			want: []bool{false, false},
		}, {
			name: "case 3",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
				queries:       [][]int{{1, 0}, {1, 2}},
			},
			want: []bool{true, true},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_closestMeetingNode(t *testing.T) {
	type args struct {
		edges []int
		node1 int
		node2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				edges: []int{2, 2, 3, -1},
				node1: 0,
				node2: 1,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				edges: []int{1, 2, -1},
				node1: 0,
				node2: 2,
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestMeetingNode(tt.args.edges, tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("closestMeetingNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

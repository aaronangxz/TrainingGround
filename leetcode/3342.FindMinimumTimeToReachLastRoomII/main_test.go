package main

import "testing"

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				moveTime: [][]int{{0, 4}, {4, 4}},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				moveTime: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}},
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				moveTime: [][]int{{0, 1}, {1, 2}},
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToReach(tt.args.moveTime); got != tt.want {
				t.Errorf("minTimeToReach() = %v, want %v", got, tt.want)
			}
		})
	}
}

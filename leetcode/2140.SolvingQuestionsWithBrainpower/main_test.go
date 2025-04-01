package main

import "testing"

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				questions: [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				questions: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			},
			want: 7,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostPoints(tt.args.questions); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

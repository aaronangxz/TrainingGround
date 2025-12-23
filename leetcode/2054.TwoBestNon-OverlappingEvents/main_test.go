package main

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			},
			want: 8,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTwoEvents(tt.args.events); got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxEvents(t *testing.T) {
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
				events: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_countTrapezoids(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			},
			want: 1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTrapezoids(tt.args.points); got != tt.want {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}

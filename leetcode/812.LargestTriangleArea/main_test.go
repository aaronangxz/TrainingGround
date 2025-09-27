package main

import "testing"

func Test_largestTriangleArea(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			},
			want: 2.00000,
		},
		{
			name: "case 2",
			args: args{
				points: [][]int{{1, 0}, {0, 0}, {0, 1}},
			},
			want: 0.50000,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTriangleArea(tt.args.points); got != tt.want {
				t.Errorf("largestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

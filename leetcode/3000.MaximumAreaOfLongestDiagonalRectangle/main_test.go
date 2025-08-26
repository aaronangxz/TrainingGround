package main

import "testing"

func Test_areaOfMaxDiagonal(t *testing.T) {
	type args struct {
		dimensions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				dimensions: [][]int{{9, 3}, {8, 6}},
			},
			want: 48,
		},
		{
			name: "case 2",
			args: args{
				dimensions: [][]int{{3, 4}, {4, 3}},
			},
			want: 12,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areaOfMaxDiagonal(tt.args.dimensions); got != tt.want {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

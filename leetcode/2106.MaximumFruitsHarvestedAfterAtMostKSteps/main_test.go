package main

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	type args struct {
		fruits   [][]int
		startPos int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				fruits:   [][]int{{2, 8}, {6, 3}, {8, 6}},
				startPos: 5,
				k:        4,
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				fruits:   [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}},
				startPos: 5,
				k:        4,
			},
			want: 14,
		},
		{
			name: "case 3",
			args: args{
				fruits:   [][]int{{0, 3}, {6, 4}, {8, 5}},
				startPos: 3,
				k:        2,
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalFruits(tt.args.fruits, tt.args.startPos, tt.args.k); got != tt.want {
				t.Errorf("maxTotalFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}

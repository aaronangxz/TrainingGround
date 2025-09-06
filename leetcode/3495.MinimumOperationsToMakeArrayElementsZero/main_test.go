package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				queries: [][]int{{1, 2}, {2, 4}},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				queries: [][]int{{2, 6}},
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.queries); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

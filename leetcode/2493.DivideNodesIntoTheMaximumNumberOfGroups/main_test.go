package main

import "testing"

func Test_magnificentSets(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnificentSets(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("magnificentSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

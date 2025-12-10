package main

import "testing"

func Test_countPermutations(t *testing.T) {
	type args struct {
		complexity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				complexity: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				complexity: []int{3, 3, 3, 4, 4, 4},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPermutations(tt.args.complexity); got != tt.want {
				t.Errorf("countPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

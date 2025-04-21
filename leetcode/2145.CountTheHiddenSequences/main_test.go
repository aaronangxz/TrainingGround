package main

import "testing"

func Test_numberOfArrays(t *testing.T) {
	type args struct {
		differences []int
		lower       int
		upper       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				differences: []int{1, -3, 4},
				lower:       1,
				upper:       6,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				differences: []int{3, -4, 5, 1, -2},
				lower:       -4,
				upper:       5,
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				differences: []int{4, -7, 2},
				lower:       3,
				upper:       6,
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArrays(tt.args.differences, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

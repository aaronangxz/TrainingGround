package main

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{3, 5, 6, 7},
				target: 9,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{3, 3, 6, 8},
				target: 10,
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{2, 3, 3, 4, 6, 7},
				target: 12,
			},
			want: 61,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums          []int
		k             int
		numOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums:          []int{1, 4, 5},
				k:             1,
				numOperations: 2,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums:          []int{5, 11, 20, 20},
				k:             5,
				numOperations: 1,
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k, tt.args.numOperations); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

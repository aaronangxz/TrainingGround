package main

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{1, 3, 5},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				arr: []int{2, 4, 6},
			},
			want: 0,
		},
		{
			name: "case 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: 16,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

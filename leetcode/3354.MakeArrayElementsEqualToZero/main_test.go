package main

import "testing"

func Test_countValidSelections(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 0, 2, 0, 3},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, 3, 4, 0, 4, 1, 0},
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidSelections(tt.args.nums); got != tt.want {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}

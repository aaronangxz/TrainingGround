package main

import "testing"

func Test_tupleSameProduct(t *testing.T) {
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
				nums: []int{2, 3, 4, 6},
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 4, 5, 10},
			},
			want: 16,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{2, 3, 4, 6, 8, 12},
			},
			want: 40,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192},
			},
			want: 1288,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.args.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

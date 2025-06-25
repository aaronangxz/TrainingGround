package main

import "testing"

func Test_kthSmallestProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{2, 5},
				nums2: []int{3, 4},
				k:     2,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{-4, -2, 0, 3},
				nums2: []int{2, 4},
				k:     6,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				nums1: []int{-2, -1, 0, 1, 2},
				nums2: []int{-3, -1, 2, 4, 5},
				k:     3,
			},
			want: -6,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestProduct(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

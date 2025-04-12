package main

import "testing"

func Test_countGoodIntegers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
				k: 5,
			},
			want: 27,
		},
		{
			name: "case 2",
			args: args{
				n: 1,
				k: 4,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				n: 5,
				k: 6,
			},
			want: 2468,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodIntegers(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("countGoodIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

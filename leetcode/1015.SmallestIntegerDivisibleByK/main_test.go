package main

import "testing"

func Test_smallestRepunitDivByK(t *testing.T) {
	type args struct {
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				K: 1,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				K: 2,
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				K: 3,
			},
			want: 3,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.args.K); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}

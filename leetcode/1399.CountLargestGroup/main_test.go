package main

import "testing"

func Test_countLargestGroup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 13,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLargestGroup(tt.args.n); got != tt.want {
				t.Errorf("countLargestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

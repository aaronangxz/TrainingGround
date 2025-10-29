package main

import "testing"

func Test_smallestNumber(t *testing.T) {
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
				n: 5,
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				n: 10,
			},
			want: 15,
		},
		{
			name: "case 3",
			args: args{
				n: 3,
			},
			want: 3,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.n); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

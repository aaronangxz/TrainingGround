package main

import "testing"

func Test_idealArrays(t *testing.T) {
	type args struct {
		n        int
		maxValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n:        2,
				maxValue: 5,
			},
			want: 10,
		},
		{
			name: "case 2",
			args: args{
				n:        5,
				maxValue: 3,
			},
			want: 11,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idealArrays(tt.args.n, tt.args.maxValue); got != tt.want {
				t.Errorf("idealArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

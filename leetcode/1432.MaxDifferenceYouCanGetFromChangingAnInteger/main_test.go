package main

import "testing"

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				num: 555,
			},
			want: 888,
		},
		{
			name: "case 2",
			args: args{
				num: 9,
			},
			want: 8,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

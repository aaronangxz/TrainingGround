package main

import "testing"

func Test_maxFreqSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "successes",
			},
			want: 6,
		},
		{
			name: "case 1",
			args: args{
				s: "aeiaeia",
			},
			want: 3,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreqSum(tt.args.s); got != tt.want {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

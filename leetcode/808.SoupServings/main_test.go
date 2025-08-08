package main

import "testing"

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				n: 50,
			},
			want: 0.62500,
		},
		{
			name: "case 2",
			args: args{
				n: 100,
			},
			want: 0.71875,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); got != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}

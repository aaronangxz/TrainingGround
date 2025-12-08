package main

import "testing"

func Test_countTriples(t *testing.T) {
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
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				n: 10,
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples(tt.args.n); got != tt.want {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}

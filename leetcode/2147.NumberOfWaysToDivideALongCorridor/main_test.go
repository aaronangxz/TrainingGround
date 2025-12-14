package main

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		corridor string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				corridor: "SSPPSPS",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				corridor: "PPSPSP",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				corridor: "S",
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.corridor); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

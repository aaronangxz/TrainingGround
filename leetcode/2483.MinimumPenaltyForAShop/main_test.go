package main

import "testing"

func Test_bestClosingTime(t *testing.T) {
	type args struct {
		customers string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				customers: "YYNY",
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				customers: "NNNNN",
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				customers: "YYYY",
			},
			want: 4,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestClosingTime(tt.args.customers); got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

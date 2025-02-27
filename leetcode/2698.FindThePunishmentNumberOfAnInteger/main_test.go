package main

import "testing"

func Test_punishmentNumber(t *testing.T) {
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
				n: 10,
			},
			want: 182,
		},
		{
			name: "case 2",
			args: args{
				n: 37,
			},
			want: 1478,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := punishmentNumber(tt.args.n); got != tt.want {
				t.Errorf("punishmentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

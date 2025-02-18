package main

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				pattern: "IIIDIDDD",
			},
			want: "123549876",
		},
		{
			name: "case 2",
			args: args{
				pattern: "DDD",
			},
			want: "4321",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.pattern); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_lengthAfterTransformations(t *testing.T) {
	type args struct {
		s string
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "abcyy",
				t: 2,
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				s: "azbk",
				t: 1,
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthAfterTransformations(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("lengthAfterTransformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

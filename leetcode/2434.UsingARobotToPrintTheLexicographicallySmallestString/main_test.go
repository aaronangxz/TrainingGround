package main

import "testing"

func Test_robotWithString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "zza",
			},
			want: "azz",
		},
		{
			name: "case 2",
			args: args{
				s: "bac",
			},
			want: "abc",
		},
		{
			name: "case 3",
			args: args{
				s: "bdda",
			},
			want: "addb",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotWithString(tt.args.s); got != tt.want {
				t.Errorf("robotWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}

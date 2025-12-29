package main

import "testing"

func Test_pyramidTransition(t *testing.T) {
	type args struct {
		bottom  string
		allowed []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				bottom:  "BCD",
				allowed: []string{"BCC", "CDE", "CEA", "FFF"},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				bottom:  "AAAA",
				allowed: []string{"AAB", "AAC", "BCD", "BBE", "DEF"},
			},
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pyramidTransition(tt.args.bottom, tt.args.allowed); got != tt.want {
				t.Errorf("pyramidTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}

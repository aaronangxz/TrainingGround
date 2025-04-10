package main

import "testing"

func Test_numberOfPowerfulInt(t *testing.T) {
	type args struct {
		start  int64
		finish int64
		limit  int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				start:  1,
				finish: 6000,
				limit:  4,
				s:      "124",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				start:  15,
				finish: 215,
				limit:  6,
				s:      "10",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				start:  1000,
				finish: 2000,
				limit:  4,
				s:      "3000",
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPowerfulInt(tt.args.start, tt.args.finish, tt.args.limit, tt.args.s); got != tt.want {
				t.Errorf("numberOfPowerfulInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

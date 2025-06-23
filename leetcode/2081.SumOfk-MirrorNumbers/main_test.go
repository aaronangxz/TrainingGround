package main

import "testing"

func Test_kMirror(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				k: 2,
				n: 5,
			},
			want: 25,
		},
		{
			name: "case 2",
			args: args{
				k: 3,
				n: 7,
			},
			want: 499,
		},
		{
			name: "case 3",
			args: args{
				k: 7,
				n: 17,
			},
			want: 20379000,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kMirror(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("kMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}

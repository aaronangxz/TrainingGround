package main

import "testing"

func Test_countBalancedPermutations(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				num: "123",
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				num: "112",
			},
			want: 1,
		},
		{
			name: "case 1",
			args: args{
				num: "12345",
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalancedPermutations(tt.args.num); got != tt.want {
				t.Errorf("countBalancedPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

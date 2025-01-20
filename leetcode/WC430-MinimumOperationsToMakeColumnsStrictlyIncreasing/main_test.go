package MinimumOperationstoMakeColumnsStrictlyIncreasing

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}},
			},
			want: 15,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.grid); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

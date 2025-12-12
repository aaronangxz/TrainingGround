package main

import (
	"reflect"
	"testing"
)

func Test_countMentions(t *testing.T) {
	type args struct {
		numberOfUsers int
		events        [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				numberOfUsers: 2,
				events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}},
			},
			want: []int{2, 2},
		},
		{
			name: "case 2",
			args: args{
				numberOfUsers: 2,
				events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}},
			},
			want: []int{2, 2},
		},
		{
			name: "case 3",
			args: args{
				numberOfUsers: 2,
				events:        [][]string{{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"}},
			},
			want: []int{0, 1},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMentions(tt.args.numberOfUsers, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countMentions() = %v, want %v", got, tt.want)
			}
		})
	}
}

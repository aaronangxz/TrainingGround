package main

import (
	"reflect"
	"testing"
)

func Test_findAllRecipes(t *testing.T) {
	type args struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				recipes:     []string{"bread"},
				ingredients: [][]string{{"yeast", "flour"}},
				supplies:    []string{"yeast", "flour", "corn"},
			},
			want: []string{"bread"},
		},
		{
			name: "case 2",
			args: args{
				recipes:     []string{"bread", "sandwich"},
				ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
				supplies:    []string{"yeast", "flour", "meat"},
			},
			want: []string{"bread", "sandwich"},
		},
		{
			name: "case 3",
			args: args{
				recipes:     []string{"bread", "sandwich", "burger"},
				ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
				supplies:    []string{"yeast", "flour", "meat"},
			},
			want: []string{"bread", "sandwich", "burger"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllRecipes(tt.args.recipes, tt.args.ingredients, tt.args.supplies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllRecipes() = %v, want %v", got, tt.want)
			}
		})
	}
}

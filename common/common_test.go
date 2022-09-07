package common

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	varA := 1
	varB := 2

	type args struct {
		a *int
		b *int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Swap",
			args{
				a: &varA,
				b: &varB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Swap(tt.args.a, tt.args.b)
			fmt.Println(varA, varB)
		})
	}
}

func TestSwapAny(t *testing.T) {
	varA := "first"
	varB := "second"
	varC := 1
	varD := 2

	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			"1.Same type string",
			args{
				a: &varA,
				b: &varB,
			},
			varA == "second" && varB == "first",
			nil,
		},
		{
			"2.Same type int",
			args{
				a: &varC,
				b: &varD,
			},
			varC == 2 && varD == 1,
			nil,
		},
		{
			"3.Different type",
			args{
				a: &varA,
				b: &varC,
			},
			varC == 1,
			ErrorInvalidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SwapAny(tt.args.a, tt.args.b); err != tt.wantErr {
				t.Errorf("SwapAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

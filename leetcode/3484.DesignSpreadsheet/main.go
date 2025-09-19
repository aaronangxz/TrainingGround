package main

import "strings"

var CELLS = make([]int32, 26*1000)

type Spreadsheet struct{}

func parseNumber(number string, j int) int {
	n := 0
	for i := j; i < len(number); i++ {
		n *= 10
		n += int(number[i] - '0')
	}
	return n
}

func parseCell(cell string) int {
	col := int(cell[0] - 'A')
	row := parseNumber(cell, 1) - 1
	return row*26 + col
}

func Constructor(rows int) Spreadsheet {
	clear(CELLS)
	return Spreadsheet{}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	CELLS[parseCell(cell)] = int32(value)
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.SetCell(cell, 0)
}

func (this *Spreadsheet) GetValue(formula string) int {
	parseVar := func(v string) int {
		if v[0] >= 'A' {
			return int(CELLS[parseCell(v)])
		} else {
			return parseNumber(v, 0)
		}
	}
	values := strings.SplitN(formula[1:], "+", 2)
	return parseVar(values[0]) + parseVar(values[1])
}

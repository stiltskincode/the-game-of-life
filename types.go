package main

import "errors"

type Matrix interface {
	Rows() int
	Cols() int
	Value(int, int) int
	SetValue(int, int, int)
}

type Grid struct {
	rows int
	cols int
	data [][]int
}

func (b *Grid) Rows() int {
	return b.rows
}

func (b *Grid) Cols() int {
	return b.cols
}

func (b *Grid) Value(x, y int) int {
	return b.data[x][y]
}

func (b *Grid) SetValue(col, row, v int) {
	b.data[col][row] = v
}

func makeGrid(rows, cols int) (*Grid, error) {
	if cols < 0 || rows < 0 {
		return nil, errors.New("Cannot initialize negative size of grid.")
	}

	data := make([][]int, cols)
	for col := 0; col < cols; col++ {
		data[col] = make([]int, rows)
	}

	g := &Grid{
		rows: rows,
		cols: cols,
		data: data,
	}
	return g, nil
}

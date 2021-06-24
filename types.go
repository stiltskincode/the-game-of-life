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

func (b *Grid) SetValue(row, col, v int) {
	b.data[row][col] = v
}

func makeGrid(rows, cols int) (*Grid, error) {
	if cols < 0 || rows < 0 {
		return nil, errors.New("Cannot initialize negative size of grid.")
	}

	data := make([][]int, rows)
	for row := 0; row < rows; row++ {
		data[row] = make([]int, cols)
	}

	g := &Grid{
		rows: rows,
		cols: cols,
		data: data,
	}
	return g, nil
}

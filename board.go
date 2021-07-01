package main

import (
	"math/rand"
)

func setRandomValues(m Matrix) {
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			m.SetValue(row, col, rand.Intn(2))
		}
	}
}

func setGlider(m Matrix) {
	centerY := m.Rows() / 2
	centerX := m.Cols() / 2

	if centerY-1 >= 0 {
		m.SetValue(centerY-1, centerX, 1)
	}

	if centerX+1 < m.Cols() {
		m.SetValue(centerY, centerX+1, 1)
	}

	if centerY+1 < m.Rows() {
		m.SetValue(centerY+1, centerX, 1)
	}

	if centerY+1 < m.Rows() {
		m.SetValue(centerY+1, centerX, 1)
	}

	if centerY+1 < m.Rows() && centerX-1 >= 0 {
		m.SetValue(centerY+1, centerX-1, 1)
	}

	if centerY+1 < m.Rows() && centerX+1 < m.Cols() {
		m.SetValue(centerY+1, centerX+1, 1)
	}

}

func getNumberOfNeighbours(y, x int, m Matrix) int {
	type point struct {
		y, x int
	}

	cols := m.Cols()
	rows := m.Rows()
	fields := map[point]int{}

	//

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			xx := (x + i + cols) % cols
			yy := (y + j + rows) % rows

			if xx != x || yy != y {
				fields[point{y: yy, x: xx}] = m.Value(yy, xx)
			}
		}
	}

	total := 0
	for _, v := range fields {
		total += v
	}

	return total
}

func calculateNextState(m Matrix) *Grid {
	rows := m.Rows()
	cols := m.Cols()
	newMatrix, _ := makeGrid(rows, cols)
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < cols; col++ {
			state := m.Value(row, col)
			neighbors := getNumberOfNeighbours(row, col, m)
			if state == 0 && neighbors == 3 {
				newMatrix.SetValue(row, col, 1)
			} else if state == 1 && (neighbors < 2 || neighbors > 3) {
				newMatrix.SetValue(row, col, 0)
			} else {
				newMatrix.SetValue(row, col, state)
			}
		}
	}
	return newMatrix
}

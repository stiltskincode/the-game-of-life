package main

import "math/rand"

func randomMatrix(m Matrix) {
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			m.SetValue(row, col, rand.Intn(2))
		}
	}
}

func getNumberOfNeighbours(y, x int, m Matrix) int {
	total := 0
	cols := m.Cols()
	rows := m.Rows()

	if y > 0 && x > 0 {
		total += m.Value(y-1, x-1)
	}

	if y > 0 {
		total += m.Value(y-1, x)
	}

	if y > 0 && x < cols-1 {
		total += m.Value(y-1, x+1)
	}

	if x < cols-1 {
		total += m.Value(y, x+1)
	}

	if y < rows-1 && x < cols-1 {
		total += m.Value(y+1, x+1)
	}

	if y < rows-1 {
		total += m.Value(y+1, x)
	}

	if y < rows-1 && x > 0 {
		total += m.Value(y+1, x-1)
	}

	if x > 0 {
		total += m.Value(y, x-1)
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

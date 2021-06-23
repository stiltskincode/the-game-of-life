package main

import "math/rand"

func randomMatrix(m Matrix) {
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			m.SetValue(col, row, rand.Intn(2))
		}
	}
}

func getNumberOfNeighbours(x, y int, m Matrix) int {
	total := 0
	cols := m.Cols()
	rows := m.Rows()

	if y > 0 && x > 0 {
		total += m.Value(x-1, y-1)
	}

	if y > 0 {
		total += m.Value(x, y-1)
	}

	if y > 0 && x < cols-1 {
		total += m.Value(x+1, y-1)
	}

	if x < cols-1 {
		total += m.Value(x+1, y)
	}

	if y < rows-1 && x < cols-1 {
		total += m.Value(x+1, y+1)
	}

	if y < rows-1 {
		total += m.Value(x, y+1)
	}

	if y < rows-1 && x > 0 {
		total += m.Value(x-1, y+1)
	}

	if x > 0 {
		total += m.Value(x-1, y)
	}

	return total
}

func calculateNextState(m Matrix) *Grid {
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

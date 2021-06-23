package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cols int = 10
const rows int = 10
const frames int = 2

func main() {
	rand.Seed(time.Now().UnixNano())

	grid, _ := makeGrid(rows, cols)
	randomMatrix(grid)

	for frame := 0; frame < frames; frame++ {
		grid = calculateNextState(grid)
	}

	for i := 0; i < grid.Cols(); i++ {
		for j := 0; j < grid.Rows(); j++ {
			fmt.Print(grid.Value(i, j))
		}
		fmt.Print("\n")
	}
}
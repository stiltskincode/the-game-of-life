package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var cols *int = flag.Int("cols", 25, "Number of columns")
var rows *int = flag.Int("rows", 25, "Number of rows")
var mapType *int = flag.Int("map", 1, "Map type: \n 1 - glider, 2 - random")

func display(m Matrix) {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			fmt.Print(m.Value(i, j))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	grid, _ := makeGrid(*rows, *cols)

	if *mapType == 1 {
		setGlider(grid)
	} else if *mapType == 2 {
		setRandomValues(grid)
	}

	for true {
		display(grid)
		grid = calculateNextState(grid)
		time.Sleep(100 * time.Millisecond)
	}
}

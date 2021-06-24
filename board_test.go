package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomMatrix(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cols := 5
	rows := 10

	firstMatrix, _ := makeGrid(rows, cols)
	randomMatrix(firstMatrix)

	secondMatrix, _ := makeGrid(rows, cols)
	randomMatrix(secondMatrix)

	diff := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if firstMatrix.Value(i, j) != secondMatrix.Value(i, j) {
				diff = true
			}
		}
	}

	if !diff {
		t.Fatalf("Matrix should have different values")
	}
}

var nextNeighboursTests = []struct {
	in  [][]int
	out []int
}{
	{[][]int{{1, 1, 1}}, []int{0, 1, 2}},
	{[][]int{{1, 1, 1}}, []int{0, 1, 2}},
	{[][]int{{1, 1, 1}}, []int{0, 0, 1}},
	{[][]int{{1}, {1}, {1}}, []int{1, 0, 2}},
}

func TestGetNumberOfNeighbours(t *testing.T) {
	for _, tt := range nextNeighboursTests {
		t.Run("Testing neighbours", func(t *testing.T) {
			cols := len(tt.in[0])
			rows := len(tt.in)

			matrix, _ := makeGrid(rows, cols)
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					matrix.SetValue(i, j, tt.in[i][j])
				}
			}

			total := getNumberOfNeighbours(tt.out[0], tt.out[1], matrix)
			if total != tt.out[2] {
				t.Fatalf("Number of number of neighbours is wrong %v", total)
			}
		})
	}
}

type moveTestCase struct {
	test   [][]int
	result [][]int
	moves  int
}

/*
 This test stage should stay in the block state
	o o o o        o o o o
	o x x o   ->   o x x o
	o x x o        o x x o
	o o o o        o o o o
*/
func state0() moveTestCase {
	test := [][]int{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	result := [][]int{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}

	tcase := moveTestCase{
		test:   test,
		result: result,
		moves:  1,
	}
	return tcase
}

/*
 Test case before and after 2 calculations
	o o o o o o       o o o o o o
	o o o o o o       o o x x o o
	o x x x x o   ->  o x o o x o
	o o o o o o       o o x x o o
	o o o o o o       o o o o o o
*/

func state1() moveTestCase {
	test := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}}
	result := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 0}, {0, 0, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0}}

	tcase := moveTestCase{
		test:   test,
		result: result,
		moves:  2,
	}
	return tcase
}

/*
 Test case before and after 2 calculations
	o o o o o o       o o o o o o
	o o x o o o       o o x x o o
	o o x x o o   ->  o x o o x o
	o o o x o o       o o x x o o
	o o o o o o       o o o o o o
*/
func state2() moveTestCase {
	test := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0}}
	result := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 0}, {0, 0, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0}}

	tcase := moveTestCase{
		test:   test,
		result: result,
		moves:  2,
	}
	return tcase
}

/*
 Test case before and after 2 calculations
	o o o o o o       o o o o o o
	o o x o o o       o o x x o o
	o o x x x o   ->  o x o o x o
	o o o o o o       o o x x o o
	o o o o o o       o o o o o o
*/
func state3() moveTestCase {
	test := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0}}
	result := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 0}, {0, 0, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0}}

	tcase := moveTestCase{
		test:   test,
		result: result,
		moves:  3,
	}
	return tcase
}

/*
 Test case before and after 4 calculations
	o o o o o o o o o     	o o o o o o o o o
	o o o o o o o o o     	o o o o o o o o o
	o o o o o o o o o   ->  o o o o o o o o o
	o o o o 1 o o o o       o o o 1 1 1 o o o
	o o o 1 1 1 o o o       o o o 1 o 1 o o o
	o o o o o o o o o		o o o 1 1 1 o o o
	o o o o o o o o o		o o o o o o o o o
	o o o o o o o o o	`	o o o o o o o o o
	o o o o o o o o o		o o o o o o o o o`
*/
func state4() moveTestCase {
	test := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	result := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	tcase := moveTestCase{
		test:   test,
		result: result,
		moves:  4,
	}
	return tcase
}

/*
	Test cases based on http://www.ibiblio.org/lifepatterns/october1970.html
*/
func TestCalculateNextState(t *testing.T) {
	tests := []moveTestCase{
		state0(),
		state1(),
		state2(),
		state3(),
		state4(),
	}

	for _, tt := range tests {
		t.Run("Testing moves", func(t *testing.T) {
			rows := len(tt.test)
			cols := len(tt.test[0])

			matrix, _ := makeGrid(rows, cols)
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					matrix.SetValue(i, j, tt.test[i][j])
				}
			}

			for i := 0; i < tt.moves; i++ {
				matrix = calculateNextState(matrix)
			}

			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					if matrix.Value(i, j) != tt.result[i][j] {
						t.Fatalf("Pattern is wrong")
					}
				}
			}
		})
	}
}

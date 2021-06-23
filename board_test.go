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

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if firstMatrix.Value(i, j) != secondMatrix.Value(i, j) {
				diff = true
			}
		}
	}

	if !diff {
		t.Fatalf("Matrix should have different values")
	}
}

func TestGetNumberOfNeighbours(t *testing.T) {
	testsMatrix := [][][]int{
		{{1, 1, 1}},
		{{1, 1, 1}}}
	testCases := [][]int{{1, 0, 2}, {0, 0, 1}}

	for i := 0; i < len(testsMatrix); i++ {
		testMatrix := testsMatrix[i]
		testCase := testCases[i]

		cols := len(testMatrix[0])
		rows := len(testMatrix)

		matrix, _ := makeGrid(rows, cols)
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				matrix.SetValue(j, i, testMatrix[i][j])
			}
		}

		total := getNumberOfNeighbours(testCase[0], testCase[1], matrix)
		if total != testCase[2] {
			t.Fatalf("Number of number of neighbours is wrong %v", total)
		}
	}
}

package main

import "testing"

func TestMakeMatrixWithNegativeValues(t *testing.T) {
	cases := [][]int{
		{-1, -1},
		{-1, 0},
		{0, -1},
	}

	for _, element := range cases {
		_, err := makeGrid(element[0], element[1])

		if err == nil {
			t.Fatalf("Expected an error but got nil")
		}
	}
}

func TestMakeMatrixWithPositiveValues(t *testing.T) {
	cases := [][]int{
		{0, 0},
		{0, 5},
		{10, 0},
		{3, 7},
	}
	for _, element := range cases {
		_, err := makeGrid(element[0], element[1])

		if err != nil {
			t.Fatalf("Expected value but got an error")
		}
	}

}

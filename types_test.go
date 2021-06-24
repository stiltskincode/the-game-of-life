package main

import "testing"

func TestMakeMatrixWithNegativeValues(t *testing.T) {

	testCases := []struct {
		name string
		in   []int
	}{
		{"Negative rows and columns", []int{-1, -1}},
		{"Negative rows but column positibe", []int{-1, 0}},
		{"Negative columns but rows positive", []int{0, -1}},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := makeGrid(tt.in[0], tt.in[1])
			if err == nil {
				t.Fatalf("Expected an error but got nil")
			}
		})
	}
}

func TestMakeMatrixWithPositiveValues(t *testing.T) {
	cases := [][]int{
		{0, 0},
		{0, 5},
		{10, 0},
		{3, 7},
	}
	for _, tt := range cases {
		t.Run("Create grid with positive elements", func(t *testing.T) {
			_, err := makeGrid(tt[0], tt[1])
			if err != nil {
				t.Fatalf("Expected value but got an error")
			}
		})
	}

}

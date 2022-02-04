package main

import (
	"testing"
)

func TestFillMatrix(t *testing.T) {

	//Arrage
	n := 2
	a = uint(2*n - 1)
	values := []uint{2, 5, 7, 1, 8, 9, 4, 7, 2}
	matrix1 := buildMatrix(a)

	//Act

	u := 0
	for k := range matrix1 {
		for j := range matrix1[k] {
			matrix1[k][j] = uint(values[u])
			u += 1
		}
	}

	//Assert
	z := 0
	for c1 := range matrix1 {
		for c2 := range matrix1[c1] {
			i := matrix1[c1][c2]
			if i != values[z] {
				t.Errorf("Incorect result. Expected %d, got %d", values[z], i)
			}
			z += 1
		}
	}
}

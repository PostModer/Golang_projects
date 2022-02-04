package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	var (
		a uint
		n uint
	)
	fmt.Print("n = ")
	fmt.Scanf("%v", &n)
	rand.Seed(time.Now().UnixNano())
	a = 2*n - 1

	var (
		size  float64 = float64(a)
		index float64 = size
		r     float64 = size
		y     int     = 0
		x     int     = int(size) - 1
	)

	//создание матрицы

	matrix := make([][]uint, a)
	for i := range matrix {
		matrix[i] = make([]uint, a)
	}
	//fmt.Print(matrix, "\n")

	// заполнение матрицы

	for k := range matrix {
		//fmt.Print("k = ", k, "\n")
		for j := range matrix[k] {
			//fmt.Print("j = ", j, "\n")
			matrix[k][j] = uint(rand.Intn(100))
		}
	}
	fmt.Print(matrix, "\n")
	answer := make([]uint, 0)
	//fmt.Print(answer, "\n")

	// задание

	for i := range matrix[0] {
		answer = append(answer, matrix[0][i])
	}
	for index < math.Pow(size, 2) {
		r -= 1
		for c := 0; c < int(r); c++ { //вертикальное перемещение вниз
			y += 1
			index += 1
			answer = append(answer, matrix[y][x])
		}
		for n := 0; n < int(r); n++ { //горизонтальное перемещение влево
			x -= 1
			index += 1
			answer = append(answer, matrix[y][x])
		}
		r -= 1
		for h := 0; h < int(r); h++ { //вертикальное перемещение вверх
			y -= 1
			index += 1
			answer = append(answer, matrix[y][x])
		}
		for o := 0; o < int(r); o++ { //горизонтальное перемещение вправо
			x += 1
			index += 1
			answer = append(answer, matrix[y][x])
		}

	}
	fmt.Print("answer = ", answer, "\n")
}

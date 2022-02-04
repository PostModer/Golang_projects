package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

//Создание Json файла
type Data struct {
	N      int
	Values []int
}

var (
	a    uint
	n    uint
	data Data
	y    int = 0
)

func (a *Data) GetN() int {
	return a.N
}

func (a *Data) GetValues() []int {
	return a.Values
}

func buildMatrix(a uint) [][]uint {
	matrix := make([][]uint, a)
	for i := range matrix {
		matrix[i] = make([]uint, a)
	}

	return matrix
}

func fillMatrix(matrix [][]uint) [][]uint {
	valData := data.GetValues()
	fmt.Print(valData)
	z := 0
	for k := range matrix {
		for j := range matrix[k] {
			matrix[k][j] = uint(valData[z])
			z += 1
		}
	}
	return matrix
}

func main() {

	d1, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	err2 := json.Unmarshal(d1, &data)
	if err2 != nil {
		panic(err)
	}

	n = uint(data.GetN())
	a = 2*n - 1
	size := float64(a)
	index := size
	r := size
	x := int(size) - 1

	//Создание матрицы

	matrix := buildMatrix(a)

	answer := make([]uint, 0)

	// заполнение матрицы

	matrix = fillMatrix(matrix)

	fmt.Print(matrix, "\n")

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

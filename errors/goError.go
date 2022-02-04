package main

import (
	"fmt"
	"os"
)

func main() {
	name, error := os.LookupEnv("NAME")
	if !error {
		panic("Переменная NAME отсутсвует")
	} else {
		fmt.Print("Name = ", name)
	}
}

package main

import (
	"fmt"
)

func main() {
	total := mult(43, 2)
	fmt.Println("total dari perkalian tersebut adalah", total)
	fmt.Println("total dari perkalian tersebut adalah", mult(21, 3))
	fmt.Println("total dari pembagian tersebut adalah", div(21, 3))
	fmt.Println("total dari pembagian tersebut adalah", div(21, mult(2, 4)))
}

func mult(a int, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("finish")
	// fmt.Println("start")
	// fmt.Println("processing")

	// num := 23
	// defer fmt.Println("result defer", num)

	// num += 7
	// num /= 3

	// // defer fmt.Println("result defer", num)
	// fmt.Println("result main:", num)

	// defer mult(2, subt(4, 8))
	// mult(4, 3)

	fmt.Println("start")
	defer loop()
	fmt.Println("finish")
	// defer mult(12, 3)
	// defer mult(1, 3)
	// defer subt(67, 12)
}

func mult(a, b int) {
	result := a * b
	fmt.Println(result)
}

func subt(a, b int) int {
	result := a - b
	fmt.Println(result)
	return result
}

func loop() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("finish looping")
}

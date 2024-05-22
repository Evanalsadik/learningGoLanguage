package main

import (
	"fmt"
)

func main() {
	var x int = 7412
	var y *int = &x

	fmt.Println("x:", x)
	fmt.Println("alamat x:", &x)

	fmt.Println("===========================")
	fmt.Println("y:", y)
	fmt.Println("alamat y:", &y)
	fmt.Println("nilai reference pointer y:", *y)

	fmt.Println("===========================")
	*y += 12
	fmt.Println("x:", x)
	fmt.Println("nilai reference pointer y:", *y)

	fmt.Println("===========================")
	increase(&x)
	fmt.Println("nilai reference pointer y:", *y)
	var w = 12
	increase(&w)
	fmt.Println("w:", w)
}

func increase(num *int) {
	*num += 6
	fmt.Println("pointer value after increased:", *num)
}

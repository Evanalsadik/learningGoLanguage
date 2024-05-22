package main

import (
	"fmt"
)

func main() {
	var x = 12
	var y = x

	fmt.Println("====================================")
	fmt.Println("x:", x)
	y += 1
	fmt.Println("y:", y)

	fmt.Println("====================================")
	fmt.Println("alamat x:", &x)
	fmt.Println("alamat y:", &y)

	fmt.Println("====================================")
	var number = 212
	decrease(number)
	fmt.Println("number:", number)
}

func decrease(num int) {
	num -= 1
	fmt.Println(num)
}

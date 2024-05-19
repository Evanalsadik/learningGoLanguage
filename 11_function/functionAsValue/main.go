package main

import (
	"fmt"
)

func main() {
	var subtract func(int, int) int = sub
	fmt.Println("hasil:", subtract(4, 2))
}

func sub(a, b int) int {
	return a - b
}

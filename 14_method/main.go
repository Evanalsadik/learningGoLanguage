package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

func main() {
	var segitiga = triangle{
		base:   21,
		height: 9,
	}

	// segitiga.area()
	hasil := segitiga.area()
	fmt.Println("luas segitiga dengan return value:", hasil)

	fmt.Println("============================================")
	fmt.Println(segitiga)
	segitiga.decreaseSize()
	fmt.Println(segitiga)
}

func (t triangle) area() float64 {
	// fmt.Println("luas segitiga:", 0.5*t.base*t.height)
	return 0.5 * t.base * t.height
}

func (t *triangle) decreaseSize() {
	t.base -= 1
	t.height -= 1
}

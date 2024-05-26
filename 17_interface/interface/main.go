package main

import (
	"fmt"
)

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Ractangle struct {
	width  float64
	length float64
}

type Square struct {
	side float64
}

// main function
func main() {
	var shape1 Shape = Square{side: 45}
	var shape2 Shape = Ractangle{width: 3, length: 5}
	var shape3 Shape = Ractangle{width: 2, length: 7}

	fmt.Printf("shape 1 : %#v \n", shape1)
	fmt.Printf("shape 2 : %#v \n", shape2)
	fmt.Printf("shape 3 : %#v \n", shape3)

	fmt.Println("=========================")
	shapes := []Shape{shape1, shape2, shape3}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Printf("%#v | memiliki luas %.1f meter persegi dan keliling %.1f meter. \n", shape, shape.getArea(), shape.getPerimeter())
	}

	fmt.Println("=========================")
	getAreaOfRactangle(Ractangle{32, 2})
	getAreaOfSquare(Square{12})
	fmt.Println("=========================")
	getArea(Ractangle{4, 5})
	getArea(Square{5})

	fmt.Println("=========================")
	// upcasting
	var square1 Shape = Square{13}
	fmt.Println("luas:", square1.getArea())
	fmt.Println("keliling:", square1.getPerimeter())
	// fmt.Println("sisi:", square1.side)

	fmt.Println("=========================")
	// downcasting
	var originalSquare Square = square1.(Square)
	fmt.Println("luas:", originalSquare.getArea())
	fmt.Println("keliling:", originalSquare.getPerimeter())
	fmt.Println("sisi:", originalSquare.side)

	fmt.Println("=========================")
	var anythings interface{}
	// var anythings any // same as empty interface above
	anythings = "damn"

	fmt.Printf("nilai dan tipe data variabel anythings: %T & %v\n", anythings, anythings)

	anythings = 523
	fmt.Printf("nilai dan tipe data variabel anythings: %T & %v\n", anythings, anythings)

	anythings = true
	fmt.Printf("nilai dan tipe data variabel anythings: %T & %v\n", anythings, anythings)

	anythings = []bool{true, false, false, true}
	fmt.Printf("nilai dan tipe data variabel anythings: %T & %v\n", anythings, anythings)

	anythings = [4]int{123, 4, 25, 2}
	fmt.Printf("nilai dan tipe data variabel anythings: %T & %v\n", anythings, anythings)
}

// method for rectangle
func (r Ractangle) getArea() float64 {
	return r.width * r.length
}

func (r Ractangle) getPerimeter() float64 {
	return 2 * (r.width * r.length)
}

// method for square
func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

// function helper satu jenis
func getAreaOfRactangle(r Ractangle) {
	fmt.Println("luas:", r.getArea())
}

func getAreaOfSquare(s Square) {
	fmt.Println("luas:", s.getArea())
}

// function helper banyak jenis
func getArea(s Shape) {
	fmt.Println("luas:", s.getArea())
}

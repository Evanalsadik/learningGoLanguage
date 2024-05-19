package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("hello from anonymous function.")
	}()

	greeting := func() {
		fmt.Println("hello from anonymous function with variable.")
	}
	greeting()

	func(nama string) {
		fmt.Printf("halo nama saya %s.\n", nama)
	}("jamal")

	sapaanDari := func(nama string) {
		fmt.Printf("halo nama saya %s.\n", nama)
	}
	sapaanDari("usrek")

	fmt.Println("=========================================")
	greeting_ := func(name string) string {
		return "hallo " + name
	}
	byeBye := func(name string) string {
		return "bye bye " + name
	}

	word("ramal.", greeting_)
	word("rusuh.", byeBye)

	fmt.Println("=========================================")

	add := func(num1, num2 int) int {
		return num1 + num2
	}
	subt := func(num1, num2 int) int {
		return num1 - num2
	}
	mult := func(num1, num2 int) int {
		return num1 * num2
	}
	div := func(num1, num2 int) int {
		return num1 / num2
	}

	fmt.Println("hasilnya adalah :", calc(12, 4, add))
	fmt.Println("hasilnya adalah :", calc(12, 4, subt))
	fmt.Println("hasilnya adalah :", calc(12, 4, mult))
	fmt.Println("hasilnya adalah :", calc(12, 4, div))
}

func word(name string, f func(name string) string) {
	fmt.Println(f(name))
}

func calc(a, b int, operator func(x, y int) int) (result int) {
	result = operator(a, b)
	return
}

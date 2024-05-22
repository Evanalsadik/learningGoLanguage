package main

import "fmt"

func main() {
	var numbers = [4]int{2, 3, 4, 3}
	var otherNumbers = numbers

	fmt.Println("====================================")
	fmt.Println("numbers :", numbers)
	fmt.Println("another numbers :", otherNumbers)

	fmt.Println("====================================")
	numbers[2] = 123
	fmt.Println("numbers :", numbers)
	fmt.Println("another numbers :", otherNumbers)

	fmt.Println("====================================")
	multBy2(otherNumbers)
	fmt.Println("another numbers :", otherNumbers)
}

func multBy2(num [4]int) {
	for i := range num {
		num[i] *= 2
	}
	fmt.Println(num)
}

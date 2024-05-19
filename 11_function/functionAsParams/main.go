package main

import (
	"fmt"
)

func main() {
	fam := []int{21, 41, 53, 2, 34, 47, 11, 46}
	hasilgp := filter(fam, isEven)
	hasilgl := filter(fam, isOdd)

	fmt.Println("hasil bilangan genap yang didapat adalah:", hasilgp)
	fmt.Println("hasil bilangan ganjil yang didapat adalah:", hasilgl)
}

func filter(numbers []int, f func(int) bool) (result []int) {
	// var result []int
	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return
}

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}

package main

import "fmt"

func main() {
	total := sum(12, 32, 34, 856, 5432, 3, 233)
	fmt.Println("total:", total)
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

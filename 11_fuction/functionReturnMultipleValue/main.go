package main

import "fmt"

func main() {
	angka := []int{21, 41, 32, 12, 343, 546, 87, 98, 2, 22}

	angkaTerkecil, angkaTerbesar, rataRata := minMax(angka)

	fmt.Printf("angka terbesar yang dimasukkan: %d\nangka terkecil yang dimasukkan: %d\ndan rata-ratanya: %d\n", angkaTerbesar, angkaTerkecil, rataRata)
}

func minMax(numbers []int) (min int, max int, avg int) {
	min = numbers[0]
	max = numbers[0]
	var temp = 0
	for _, number := range numbers {
		temp += number

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	avg = temp / len(numbers)

	return
}

package main

import (
	"fmt"
)

func main() {
	// FORMAT ANGKA
	var angka float32

	fmt.Print("masukkan angka desimal: ")
	fmt.Scan(&angka)

	fmt.Printf("angka yang anda masukkan adalah: %.2f\n", angka)
}

package main

import (
	"fmt"
	"strconv"
)

var names []string

func main() {
	fmt.Println("names :", names)
	var input string
	fmt.Print("masukkan jumlah nama: ")
	fmt.Scanln(&input)

	len, _ := strconv.Atoi(input)
	for i := 0; i < len; i++ {
		fmt.Println("masukkan nama: ")
		fmt.Scanln(&input)
		names = append(names, input)
	}
	fmt.Println("names :", names)
}

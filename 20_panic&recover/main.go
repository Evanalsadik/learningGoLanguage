package main

import "fmt"

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("terjadi panic:", r)
	// 	}
	// }()

	var input string
	fmt.Print("masukkan nama: ")
	fmt.Scanln(&input)

	fmt.Println("start")
	if !isEmpty(input) {
		fmt.Println(input)
	}
	fmt.Println("done.")
}

func isEmpty(input string) (empty bool) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("terjadi panic:", r)
		}
		empty = true
	}()

	if input == "" {
		panic("input tidak boleh kosong.")
	}
	empty = false
	return
}

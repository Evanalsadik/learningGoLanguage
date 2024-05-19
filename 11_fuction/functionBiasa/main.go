package main

import (
	"fmt"
)

var age int = 18

func main() {
	fmt.Println("================================= void function")
	helloWorld()
	fmt.Printf("main function age %d...\n", age)
	fmt.Println("================================= function with params")
	nameReturn("damar", 43)
	nameReturn("usrek", 49)
	nameReturn("rek", 53)
	fmt.Println("================================= function with two params")
	addition(23, 9)
}

func helloWorld() { //void function
	var name string = "jamal" //local variable
	var age = 20
	fmt.Printf("function hello %s...\n", name)
	fmt.Printf("helloWorld function age %d...\n", age)
}

func nameReturn(username string, myage int) {
	fmt.Printf("perkenalkan nama saya %s umur saya %d tahun.\n", username, myage)
}

func addition(firstNumber int, secondNumber int) {
	fmt.Println("hasil penjumlahan:", firstNumber+secondNumber)
}

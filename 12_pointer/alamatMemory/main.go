package main

import (
	"fmt"
)

func main() {
	var name string = "jamal"
	var age int = 18

	fmt.Println("====================================")
	fmt.Println("nama:", name)
	fmt.Println("alamat nama:", &name)

	fmt.Println("====================================")
	fmt.Println("umur:", age)
	fmt.Println("alamat umur:", &age)
}

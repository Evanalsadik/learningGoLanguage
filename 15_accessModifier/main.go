package main

import (
	"fmt"
	"golang-introduction/15_accessModifier/library"
)

func main() {
	fmt.Println("second in a minute:", library.HourInDay)
	fmt.Println("Node:", library.Node)

	fmt.Println("===================")
	hasil := library.DayToHour(2)
	fmt.Printf("dua hari sama dengan %d jam", hasil)
}

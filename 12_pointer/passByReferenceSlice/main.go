package main

import (
	"fmt"
)

func main() {
	var numbers = []int{21, 312, 4, 32, 3, 54, 2}
	var pNum = numbers

	fmt.Println("========================")
	fmt.Println(numbers)
	fmt.Println(pNum)

	fmt.Println("========================")
	numbers[2] = 100
	fmt.Println(numbers)
	fmt.Println(pNum)

	fmt.Println("========================")
	multBy2(numbers)
	fmt.Println(numbers)
	fmt.Println(pNum)
}

func multBy2(nums []int) {
	for num := range nums {
		nums[num] *= 2
	}
	fmt.Println(nums)
}

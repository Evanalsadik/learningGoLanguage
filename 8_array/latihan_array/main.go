package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	// Membaca kapasitas array
	capacityInput, _ := scanner.ReadString('\n')
	capacityInput = strings.TrimSpace(capacityInput)
	capacity, _ := strconv.Atoi(capacityInput)
	
	// Membaca elemen-elemen arra
	elementsInput, _ := scanner.ReadString('\n')
	elementsInput = strings.TrimSpace(elementsInput)
	elements := strings.Split(elementsInput, " ")
	
	// Membuat array dengan kapasitas yang diberikan
	array := make([]int, capacity)

	// Mengisi array dengan elemen-elemen yang diberikan
	for i, element := range elements {
		element, _ := strconv.Atoi(element)
		array[i] = element
	}

	// Mencetak angka-angka genap
	for _, element := range array {
		if element % 2 == 0 {
			fmt.Println(element)
		}
	}
}


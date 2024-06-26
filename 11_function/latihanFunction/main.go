package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Masukan banyak bamboo: ")
	scanner.Scan()
	jumlahBamboo, _ := strconv.Atoi(scanner.Text())

	jumlahSekat := make([]int, jumlahBamboo)
	for i := 0; i < jumlahBamboo; i++ {
		// fmt.Printf("Sekat bamboo ke-%d: ", i+1)
		scanner.Scan()
		jumlah, _ := strconv.Atoi(scanner.Text())
		jumlahSekat[i] = jumlah
	}
	// fmt.Print("Berapa kali potong? ")
	scanner.Scan()
	jumlahPotong, _ := strconv.Atoi(scanner.Text())

	cutOffTheBamboo(jumlahSekat, jumlahPotong)
}

func cutOffTheBamboo(jumlahSekat []int, jumlahPotong int) {
	for i := 1; i <= jumlahPotong; i++ {
		fmt.Printf("Potongan ke- %d\n", i)
		for j := 0; j < len(jumlahSekat); j++ {
			if jumlahSekat[j] > 0 {
				jumlahSekat[j] -= 1
			}
			fmt.Println(jumlahSekat[j])
		}
	}
	fmt.Println()
}


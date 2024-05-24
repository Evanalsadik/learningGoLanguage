package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Input
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	marble, _ := strconv.Atoi(scanner.Text())

	// Call function
	giveMarble(&giver, &receiver, marble)
	fmt.Printf("%d \n", giver)
	fmt.Printf("%d \n", receiver)
}

func giveMarble(pemberiKelereng, penerimaKelereng *int, jumlahKelereng int) {
	if *pemberiKelereng < jumlahKelereng {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
	}
	if *pemberiKelereng > jumlahKelereng {
		*pemberiKelereng -= jumlahKelereng

		*penerimaKelereng += jumlahKelereng
	}

}

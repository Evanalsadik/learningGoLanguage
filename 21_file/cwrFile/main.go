package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "/home/left/Desktop/learningGoLanguage/21_file/data/"
var fileName = "example.txt"
var filePath = path + fileName

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan nama: ")
	scanner.Scan()
	input = scanner.Text()

	createFile()
	writeFile(input)
	readFile()
}

func createFile() {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println("file", fileName, "berhasil dibuat.")
	}
}

func writeFile(input string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writter := bufio.NewWriter(file)
	writter.WriteString(input + "\n")
	writter.Flush()
}

func readFile() {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

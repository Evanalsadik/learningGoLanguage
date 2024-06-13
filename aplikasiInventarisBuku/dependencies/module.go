package dependencies

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

var Path = "/home/left/Desktop/learningGoLanguage/aplikasiInventarisBuku/data/"
var FileName = "daftar-buku.csv"
var FilePath = Path + FileName

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

func ContainEmptyValue(cekNilai Book) bool {
	return cekNilai.Id == 0 || cekNilai.Pages == 0 || cekNilai.Author == "" || cekNilai.Title == "" || cekNilai.ReleaseYear == ""
}

func CreateFile() {
	// membuat file csv
	_, err := os.Stat(FilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(FilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// membuatkan header
		writer := csv.NewWriter(file)
		defer writer.Flush()

		header := []string{"Id", "Title", "Author", "ReleaseYear", "Pages"}
		err = writer.Write(header)
		if err != nil {
			fmt.Println("Error writing header to CSV:", err)
			return
		}
		fmt.Println("file", FileName, "successfully created.")
	}
}

func InputUpdatedValue(id int) Book {
	var updatedValue Book
	scanner := bufio.NewScanner(os.Stdin)
	// cannot update Id
	updatedValue.Id = id

	fmt.Println("Silahkan isi informasi buku dibawah ini.")
	fmt.Print("Updated Title: ")
	scanner.Scan()
	updatedValue.Title = scanner.Text()
	fmt.Print("Updated Author: ")
	scanner.Scan()
	updatedValue.Author = scanner.Text()
	fmt.Print("Updated ReleaseYear: ")
	scanner.Scan()
	updatedValue.ReleaseYear = scanner.Text()
	fmt.Print("Updated Pages: ")
	scanner.Scan()
	updatedValue.Pages, _ = strconv.Atoi(scanner.Text())

	return updatedValue
}
func SaveUpdatedDataToCSV(books []Book) {
	file, err := os.OpenFile(FilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	// header := []string{"Id", "Title", "Author", "ReleaseYear", "Pages"}
	// if err := writer.Write(header); err != nil {
	// 	fmt.Println("Error writing header to CSV:", err)
	// 	return
	// }

	// Write records
	for _, book := range books {
		record := []string{
			strconv.Itoa(book.Id),
			book.Title,
			book.Author,
			book.ReleaseYear,
			strconv.Itoa(book.Pages),
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
		}
	}
}

func LoadDataFromCSV() []Book {
	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var temp []Book

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading record:", err)
			return nil
		}

		var book Book
		book.Id, _ = strconv.Atoi(record[0])
		book.Title = record[1]
		book.Author = record[2]
		book.ReleaseYear = record[3]
		book.Pages, _ = strconv.Atoi(record[4])

		temp = append(temp, book)
	}
	return temp
}

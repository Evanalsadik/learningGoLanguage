package dependencies

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ViewAllBooks() {
	temp := LoadDataFromCSV()
	for i, b := range temp {
		fmt.Println("====================================")
		fmt.Printf("Book - %d \nBook Id: [%d]\nBook Title: [%s]\nBook Author: [%s]\nBook ReleaseYear: [%s]\nBook Pages: [%d]\n",
			i+1, b.Id, b.Title, b.Author, b.ReleaseYear, b.Pages)
		fmt.Printf("====================================\n\n")
	}
}

func AddNewBook() []Book {
	var inputBuku Book
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Silahkan isi informasi buku dibawah ini.")
	fmt.Print("Id: ")
	scanner.Scan()
	inputBuku.Id, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Title: ")
	scanner.Scan()
	inputBuku.Title = scanner.Text()
	fmt.Print("Author: ")
	scanner.Scan()
	inputBuku.Author = scanner.Text()
	fmt.Print("ReleaseYear: ")
	scanner.Scan()
	inputBuku.ReleaseYear = scanner.Text()
	fmt.Print("Pages: ")
	scanner.Scan()
	inputBuku.Pages, _ = strconv.Atoi(scanner.Text())

	if ContainEmptyValue(inputBuku) {
		fmt.Println("Invalid input data.")
		return []Book{}
	}
	return []Book{inputBuku}
}

func SaveDataToCSV(datas []Book) {
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write records
	for _, data := range datas {
		record := []string{
			strconv.Itoa(data.Id),
			data.Title,
			data.Author,
			data.ReleaseYear,
			strconv.Itoa(data.Pages),
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
		}
	}
}

func UpdateBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan Id: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	books := LoadDataFromCSV()
	var found bool

	for i, book := range books {
		if id == book.Id {
			result := InputUpdatedValue(id)
			if ContainEmptyValue(result) {
				fmt.Println("Invalid input data.")
				return
			}
			books[i] = result
			found = true
			break
		}
	}
	if !found {
		fmt.Println("The ID you entered is not listed in the CSV file.")
		return
	}

	SaveUpdatedDataToCSV(books)
	fmt.Println("Book updated successfully.")
}

func DeleteBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan Id: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	books := LoadDataFromCSV()
	var (
		found       bool
		result      []Book
		deletedData Book
	)

	for _, book := range books {
		if id != book.Id {
			result = append(result, book)
		} else {
			found = true
			deletedData = book
		}
	}

	if !found {
		fmt.Println("The ID you entered is not listed in the CSV file.")
		return
	}

	fmt.Println("data yang berisikan", deletedData, "berhasil dihapus.")
	SaveUpdatedDataToCSV(result)

}

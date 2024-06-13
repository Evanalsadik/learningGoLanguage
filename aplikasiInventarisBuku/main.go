package main

import (
	"bufio"
	"fmt"
	"golang-introduction/aplikasiInventarisBuku/dependencies"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("===============================================")
	fmt.Println("|========== APLIKASI INVENTARIS GO ===========|")

	var (
		condition bool
		semuaBuku []dependencies.Book
	)
	condition = true
	for condition {
		fmt.Printf("Silahkan pilih menu anda:\n\t 1. View all books\n\t 2. Add new book\n\t 3. Update book\n\t 4. Delete book\n\t 5. Exit\n\t 0. Create file\n")
		fmt.Print("pilihan anda: ")
		scanner.Scan()
		pilihan, _ := strconv.Atoi(scanner.Text())

		switch pilihan {
		case 1:
			{
				fmt.Println("Anda melilih untuk melihat daftar buku.")
				dependencies.ViewAllBooks()
				fmt.Println("selesai.")
				fmt.Println("===============================================")
			}
		case 2:
			{
				fmt.Println("Anda melilih untuk menambahkan daftar buku.")
				semuaBuku = dependencies.AddNewBook()
				dependencies.SaveDataToCSV(semuaBuku)
				fmt.Println("selesai.")
				fmt.Println("===============================================")
			}
		case 3:
			{
				fmt.Println("Anda melilih untuk merubah daftar buku.")
				dependencies.UpdateBook()
				fmt.Println("selesai.")
				fmt.Println("===============================================")
			}
		case 4:
			{
				fmt.Println("Anda melilih untuk menghapus daftar buku.")
				dependencies.DeleteBook()
				fmt.Println("selesai.")
				fmt.Println("===============================================")
			}
		case 5:
			{
				condition = false
				fmt.Println("bye.")
			}
		case 0:
			{
				fmt.Println("Anda melilih untuk membuat file baru.")
				dependencies.CreateFile()
				fmt.Println("selesai.")
				fmt.Println("===============================================")
			}
		default:
			{
				condition = true
				fmt.Println("===============================================")
				fmt.Println("pilih dengan nomor menu yang telah disediakan.")
			}
		}

	}
}

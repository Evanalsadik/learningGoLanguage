package main

import (
	"fmt"
)

type Student struct {
	// Isi struct ini
	nama string
	umur int
}

func main() {
	// Tulis kode disini

	var student1 Student
	student1.nama = "Sammy"
	student1.umur = 17
	student1.Introduction()
}

func (s *Student) Introduction() {
	// Tulis kode disini
	fmt.Printf("Hello, my name is %s. Im %d years old\n", s.nama, s.umur)
}

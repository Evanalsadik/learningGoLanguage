package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main(){
  //package fmt
  fmt.Println("=====================================")
  var name string
  var age int64

  fmt.Print("enter random name: ")
  fmt.Scanln(&name)
  fmt.Print("enter random age: ")
  fmt.Scanln(&age)

  birtday := 2024 - age

  fmt.Printf("hi! my name is %s, i was born in %d so i'm %d years old now.\n", name, birtday, age)

  //package bufio
  fmt.Println("=====================================")
  var fullName string
  var message string
  var id int

  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("todo list alakadarnya")

  fmt.Print("silahkan masukkan nama lenngkap anda: ")
  scanner.Scan()
  fullName = scanner.Text()
  
  fmt.Print("silahkan masukkan nomor id: ")
  scanner.Scan()
  id, _ = strconv.Atoi(scanner.Text())

  fmt.Print("tuliskan pesan anda: ")
  scanner.Scan()
  message = scanner.Text()

  fmt.Printf("_ Pemilik catatan\t: %s\n_ Id\t\t\t: %d\n_ Todo\t\t\t: %s\n", fullName, id, message)

}

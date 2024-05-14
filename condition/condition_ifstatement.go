package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  fmt.Println("===================================")
  fmt.Println("satu kondisi")
  if true {
    fmt.Println("...")
  }
  fmt.Println("done.")

  fmt.Println("===================================")
  fmt.Println("dua kondisi")
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print("masukkan angka: ")
  scanner.Scan()
  
  /* jika variabel angka disimpan didepan 'if'/dalam 'if statement', maka variabel
  tersebut akan bersifat lokal (hanya bisa diakses didalam 'if statement' 
  tersebut.) */
  if angka, _ := strconv.Atoi(scanner.Text()); angka >12 && angka <20 {
    fmt.Println("anda masuk himpunan.")
  } else {
    fmt.Println("anda bukan kawan saya.")
  }
  
  fmt.Println("===================================")
  fmt.Println("tiga/lebih kondisi")

  var suhu int
  var kondisiTubuh string

  fmt.Print("masukkan suhu tubuh: ")
  scanner.Scan()
  suhu, _ = strconv.Atoi(scanner.Text())

  if suhu >= 37 && suhu < 41 {
    kondisiTubuh = "anda sakit keras, butuh es batu."
  } else if suhu >= 35 && suhu < 37 {
    kondisiTubuh = "anda lumayan sehat, dikit."
  } else if suhu >= 29 && suhu < 35 {
    kondisiTubuh = "anda kedinginan, sheesh..."
  } else {
    kondisiTubuh = "mungkin anda sudah mati."
  }
  fmt.Println(kondisiTubuh)

  fmt.Println("===================================")
  fmt.Println("nested condition/kondisi bersarang")
  
  x := -17
  y := -16

  if x > 0 {
    if y > 0 {
      fmt.Println(x, "dan", y, "adalah bilangan positif.")
    } else {
      fmt.Println(x, "adalah bilangan positif dan", y, "adalah bilangan negatif.")
    }
  } else if y > 0 {
    fmt.Println(x, "adalah bilangan negatif dan", y, "adalah bilangan positif.")
  } else {
    fmt.Println(x, "dan", y, "adalah bilangan negatif.")
  }

}

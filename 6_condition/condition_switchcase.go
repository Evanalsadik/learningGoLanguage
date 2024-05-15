package main

import (
  "fmt"
)

func main() {
  //bentuk pertama
  fmt.Println("=============================================")
  var poin int8 = 5

  switch poin {
  case 8: {
    fmt.Println("bagus.")
  } 
  case 7, 6, 5: {
    fmt.Println("lumayan.")
  }
  case 4: {
    fmt.Println("payah.")
  }
  default: {
    fmt.Println("sudahlah.")
  }
  }

  //bentuk kedua
  fmt.Println("=============================================")
  switch {
  case poin > 8: {
    fmt.Println("bagus 2.")
  }
  case poin < 8 && poin > 4: {
    fmt.Println("mengecewakan.")
  }
  fallthrough // <- mau bagaimanapun kondisinya, jika ada 'fallthrough', maka case dibawahnya akan dieksekusi.
  default: {
    fmt.Println("terparah sepanjang masa.")
  }
    
  }
}

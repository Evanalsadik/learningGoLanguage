package main

import (
  "fmt"
)

func main() {
  //standard
  fmt.Println("=======================")
  for i := 1; i <= 5; i++ {
    fmt.Println("\t• angka", i)
  }

  //variabel 'init' terpisah
  fmt.Println("=======================")
  var i = 7
  for i <= 15 {
    fmt.Println("\t• angka", i)
    i++
  }

  //menggentikan looping menggunakan 'if condition'
  fmt.Println("=======================")
  var index = 0
  for {
    fmt.Println("\t• angka", index)
    index++
    if index == 10 {
      break
    }
  }
  
  //nested loop
  fmt.Println("=======================")
  for x := 1; x <= 5; x++ {
    for y := 1; y <= 5; y++ {
      fmt.Print(y, " ")
    }
    fmt.Println()
  }
}

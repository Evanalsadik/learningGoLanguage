package main

import (
  "fmt"
)

func main() {
  //break
  fmt.Println("=======================")
  fmt.Print(" •). ")
  for i := 1; i <= 10; i++ {
    if i == 8 {
      break
    }
    fmt.Print(i, " ")
  }
  fmt.Println("\n     done.\n")

  fmt.Println("=======================")

  //continue
  fmt.Println("=======================")
  fmt.Print(" •). ")
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      continue
    }
    fmt.Print(i, " ")
  }
  fmt.Println("\n     done.\n")

  // nesdted loop 1
  fmt.Println("=======================")
  for x := 1; x <= 3; x++ {
    for y := 1; x <= 5; y++ {
      if y == 4 {
        break
        // continue
      }
      fmt.Print(y, " ")
    }
    fmt.Println()
  }

  // nesdted loop 2
  fmt.Println("=======================")
  var poin bool = false
  for i := 0; i < 3; i++ {
    if poin {
      break
    }
    for j := 0; j < 3; j++ {
      if j == 12 {
        poin = true // Mengatur variabel boolean untuk menghentikan loop terluar
        break
        // continue
      }
      fmt.Println(i, j)
    }
  }

  fmt.Println("\ndone.")

}

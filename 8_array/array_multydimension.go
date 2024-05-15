package main

import (
  "fmt"
)

func main() {
  fmt.Println("=====================================")
  var matrix = [2][3]int {
    {5, 82, 638},
    {72, 93, 61},
  }

  //array 2 dimensi
  fmt.Println(matrix)

  //memisahkan array 2 dimensi menjadi 1 dimensi
  fmt.Println(matrix[0])
  fmt.Println(matrix[1])

  //mengambil satu elemen dari array 2 dimensi
  fmt.Println(matrix[1][2])


}

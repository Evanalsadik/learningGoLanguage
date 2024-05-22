package main

import (
  "fmt"
)

func main() {
  fmt.Println("=========================")
  //membuat array
  var arr = [3]int {1,2,3}
  fmt.Println(arr)

  //memanggil array secara spesifik
  fmt.Println(arr[1])

  //reassign array
  arr[2] = 100
  fmt.Println(arr[2])

  fmt.Println("=========================")
  //membuat array (cara lain)
  var name [4]string
  name[0] = "qubic"
  name[1] = "quatro"
  name[2] = "qwebec"
  name[3] = "square"
  fmt.Println(name)

  fmt.Println("=========================")
  //membuat array (cara lain lagi)
  var cars = [...]string {"bmw", "ferrari", "mclaren", "mcqueen"}
  fmt.Println(cars)
  fmt.Println("jumlah mobil yang ada: ", len(cars))

  fmt.Println("=========================")
  //array looping
  for i := 0; i < len(cars); i++ {
    fmt.Printf("mobil ke-%d adalah %s\n", i+1, cars[i])
  }

  fmt.Println("=========================")
  //array looping (dengan for range)
  for i, car := range cars {
    fmt.Printf("mobil ke-%d adalah %s\n", i+1, car)
  }

  fmt.Println("=========================")
  //array looping (tanpa index)
  fmt.Printf("daftar mobil yang tersedia: \n")
  for _, car := range cars {
    fmt.Printf("\t• %s\n", car)
  }

  fmt.Println("=========================")
  //array looping (ntah apa gunanya, tapi mungkin bisa berguna)
  fmt.Printf("hitung mobil yang ada: \n")
  for i := range cars {
    fmt.Printf("mobil ke-%d\n", i+1)
  }

  fmt.Println("=========================================")
  //looping, condition, dan array jika digabung
  var genap = [10]int {52, 73, 62, 93, 536672, 5499, 7162, 5273, 53000, 19}

  for _, angkaGenap := range genap {
    if angkaGenap % 2 == 0 {
      fmt.Printf("angka %d merupakan angka genap.\n", angkaGenap)
    } else {
      fmt.Printf("angka %d merupakan angka ganjil.\n", angkaGenap)
    }
  }

}

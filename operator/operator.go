package main

import "fmt"

func main(){
  x := 93
  y := 76

  fmt.Println("\n====================================")
  fmt.Println("ARITMATIKA")
  fmt.Println("1). hasil operasi pertambahan: ", x + y)
  fmt.Println("2). hasil operasi pengurangan: ", x - y)
  fmt.Println("3). hasil operasi perkalian: ", x * y)
  fmt.Println("4). hasil operasi pembagian: ", x / y)
  fmt.Println("5). hasil operasi modulo: ", x % y)

  var a int = 836
  var b int

  fmt.Println("\n====================================")
  fmt.Println("ASSIGNMENT")
  b = a
  fmt.Println("1). nilai b : ", b)
  b += 77
  fmt.Println("2). nilai b (pertambahan) : ", b)
  b -= 77
  fmt.Println("3). nilai b (pengurangan) : ", b)
  b *= 77
  fmt.Println("4). nilai b (perkalian) : ", b)
  b /= 77
  fmt.Println("5). nilai b (pembagian) : ", b)
  b %= 77
  fmt.Println("6). nilai b (modulo) : ", b)

  m := 6487
  n := 863

  fmt.Println("\n====================================")
  fmt.Println("COMPARISON")
  fmt.Println("1). m setara dengan n : ", m == n)
  fmt.Println("2). m tidak setara dengan n : ", m != n)
  fmt.Println("3). m lebih besar/setara dengan n : ", m >= n)
  fmt.Println("4). m lebih kecil/setara dengan n : ", m <= n)
  fmt.Println("5). m lebih besar dengan n : ", m > n)
  fmt.Println("6). m lebih kecil dengan n : ", m < n)

  t := true
  f := false

  fmt.Println("\n====================================")
  fmt.Println("LOGIC")
  fmt.Println("1). apakah keduanya benar : ", t && f)
  fmt.Println("2). apakah salah satunya benar : ", t || f)
  fmt.Println("3). kebalikan/negasi dari t : ", !t)
  fmt.Println("3). kebalikan/negasi dari f : ", !f)
  fmt.Println("prioritas : ", false || false && true)

}

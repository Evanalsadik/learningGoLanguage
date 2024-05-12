package main

import "fmt"

func main(){
  // cara pertama
  var name string
  name = "leftside"

  // cara kedua
  var age string = "7 years old.\n"

  // cara ketiga
  var (
    candy int
    addres string
  )

  candy = 19
  addres = "Bandung"

  //cara keempat
  var (baloon, uang = 5, "dua ribu")


  // MANIFEST TYPING
  fmt.Println("=======================================================")
  fmt.Println("MANIFEST TYPING")
  fmt.Println("hello, my name is", name, "and i'm", age)
  fmt.Printf("hello, my name is %s and i'm %s \n and i hate kiddos.\nby the way, i'll gave you %d candies if you want.\njust come to my house at %s.\n\n", name, age, candy, addres)
  fmt.Printf("aku ingin membeli %d balon. tapi aku hanya punya uang %s rupiah saja.\n\n", baloon, uang)

  //cara pertama
  day := "Friday"
  date := "13"
  month := "january"
  
  //cara kedua (variabel konstan)
  const phi = 3.14
  
  //cara ketiga (variabel mati)
  planet, _ := "saturnus", "matahari"

  // TYPE INFERENCE
  fmt.Println("=======================================================")
  fmt.Println("TYPE INFERENCE")
  fmt.Println(day, date, month, " 2014.")
  fmt.Println(phi)
  fmt.Println(planet)
}

package main

import "fmt"

func main(){
  fmt.Println("==========================================")
  //Numeric (INTEGER)
  var positiveNumber uint8 = 69
  fmt.Printf("bilangan bulat positif (khusus 'uint'): %d\n", positiveNumber)
  
  var integer1 int8 = 69
  var integer2 int8 = -69
  fmt.Printf("bilangan bulat (positif/negatif 'int'): %d / %d\n", integer1, integer2)
  
  fmt.Println("==========================================")
  //Numeric (INTEGER)
  var decimalNumber = 3.14
  fmt.Printf("bilangan desimal: %f\n", decimalNumber)
  fmt.Printf("bilangan desimal (terbatas): %.2f\n", decimalNumber)

  fmt.Println("==========================================")
  //Boolean
  var exist = false
  fmt.Printf("variabel exist bernilai: %t\n", exist)

  fmt.Println("==========================================")
  //String
    // geberal (""/'')
  var message1 string = "halo selamat pagi."
  var message2 string = "mulai hari anda dengan bekerja di sini."
  fmt.Printf("anda mendapat pesan berisi ucapan sebagai berikut: \n\t• %s\n\t• %s\n\n", message1, message2)
    // string literal
  var literalMessage string = `
        • selamat, hari ini anda dipecat.
        • silahkan keluar.`
  fmt.Printf("anda mendapat pesan berisi ucapan sebagai berikut: %s\n\n", literalMessage)
    // string concatenation
  var firstWord string = "hanya"
  var secondWord string = "bercanda"
  fmt.Println("anda mendapat pesan berisi ucapan sebagai berikut:")
  fmt.Println("\t" + "• maaf " + firstWord + " " + secondWord + ".\n")

    // len function
  var word string = "apalah yang penting ada."
  fmt.Println("jumlah karakter dalam variabel word adalah:", len(word))

}

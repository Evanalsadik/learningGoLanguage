package main

import (
  "fmt"
)

func main() {
  fmt.Println("===================================")
  numberSlice := []int {2, 4, 6, 8}
  fmt.Println(numberSlice)
  
  fmt.Println("panjang slice:", len(numberSlice))
  fmt.Println("kapasitas slice:", cap(numberSlice))
  
  fmt.Println("===================================")
  daftarKota := make([]string, 5, 10)
  daftarKota[0] = "Bandung"
  daftarKota[1] = "Jakarta"
  daftarKota[2] = "Palembang"
  daftarKota[3] = "Padang"
  daftarKota[4] = "Medan"
  // daftarKota[5] = "aceh" // error out of bounds

  fmt.Println("panjang slice:", len(daftarKota))
  fmt.Println("kapasitas slice:", cap(daftarKota))
  fmt.Println(daftarKota)
  

  fmt.Println("===================================")
  planet := [3]string {"jupiter", "bumi", "venus"}
  fmt.Println("planet :", planet)
  // planet[3] = "pluto" // error out of bounds
  
  fmt.Println("===================================")
  negara := make([]string, 5, 10)
  negara[0] = "Denmark"
  negara[1] = "Bekasi"
  negara[2] = "Mesir"
  negara[3] = "Russia"
  negara[4] = "Jepang"
  
  fmt.Println("negara :", negara)
  fmt.Println("panjang slice:", len(negara))
  
  // cara menambahkan elemen kedalam slice negara yang sudah penuh panjangnya
  negara = append(negara, "Bogor")
  fmt.Println("negara :", negara)
  fmt.Println("panjang slice:", len(negara))
  fmt.Println("kapasitas slice:", cap(negara))


  fmt.Println("===================================")
  //
  ages := [4]int {27, 48, 72, 91}
 /*  memotong array ages dari indek ke 0 sebanyak 2 elemen (index 0 dan 1)
  jika value pada slice variabel berubah, maka value pada root 
  variabel akan ikut berubah juga
  dan jika melakukan 'append' pada slice variabel, maka elemen pada
  index terakhir dari root variabel akan berubah(tidak ikut bertambah,
  hanya berubah) */
  sliceAges := ages[0:2]
  fmt.Println("elemen yang diambil: ", sliceAges)

}

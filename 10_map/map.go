package main

import (
  "fmt"
)

func main() {
  fmt.Println("===================================")
  //cara membuat map, gunakan keyword 'map[key]value' untuk key dan value isi dengab tipe datanya
  kota := map[string]string{
    "amerika":"london",
    "russia":"mumbai",
    "afrika selatan":"tokyo",
  }

  fmt.Println(kota)
  

  fmt.Println("===================================")
  var negara = make(map[string]string)
  negara["london"] = "prancis"
  negara["berlin"] = "belgia"
  negara["jakarta"] = "italia"
  
  fmt.Println(negara)
  fmt.Println("london adalah ibukota negara", negara["london"])
  

  fmt.Println("===================================")
  negara["london"] = "belanda"
  fmt.Println("london adalah ibukota negara", negara["london"])
  

  fmt.Println("===================================")
  delete(negara, "london")
  fmt.Println("london adalah ibukota negara", negara["london"])
  

  fmt.Println("===================================")
  names := map[int]string {
    1:"damar",
    2:"jamal",
    3:"kamal",
    4:"usrek",
  }
  for key, value := range names {
    fmt.Println(key, ":", value)
  }
}

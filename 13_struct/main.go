package main

import (
	"fmt"
)

type buah struct {
	nama  string
	hijau bool
}

type fighterJet struct {
	merek  string
	warna  string
	jumlah int
	jenis  string
	buah
}

func main() {
	var raptor fighterJet
	raptor.merek = "F-22"
	raptor.warna = "abu-abu"
	raptor.jumlah = 12
	raptor.jenis = "stealth"

	fmt.Println("raptor:", raptor)
	fmt.Println("merek:", raptor.merek)
	fmt.Println("warna:", raptor.warna)
	fmt.Println("jumlah:", raptor.jumlah)
	fmt.Println("jenis:", raptor.jenis)

	fmt.Println("====================================")
	var chengdu = fighterJet{}
	chengdu.merek = "J-20"
	chengdu.warna = "abu-abu"
	chengdu.jumlah = 3
	chengdu.jenis = "hybrid"

	fmt.Println("chengdu:", chengdu)

	var rafale = fighterJet{"_", "abu-abu", 789, "boomber", buah{"pisang", false}}

	fmt.Println("rafale:", rafale)

	var regularFIghterJet = fighterJet{
		warna:  "hitam",
		jumlah: 67876,
		jenis:  "fighter",
		merek:  "F-16",
	}

	fmt.Println("blyat:", regularFIghterJet)

	fmt.Println("====================================")
	var y = regularFIghterJet
	fmt.Println("x:", regularFIghterJet)
	fmt.Printf("alamat x: %p\n", &regularFIghterJet)
	fmt.Println("y:", y)
	fmt.Printf("alamat y: %p\n", &y)

	fmt.Println("====================================")
	y.jenis = "boommber"
	fmt.Println(regularFIghterJet)
	fmt.Println(y)

	fmt.Println("====================================")
	var x = fighterJet{merek: "lilin", warna: "hijau", jenis: "padat", jumlah: 67}
	var xi *fighterJet = &x
	fmt.Printf("alamat xi: %p\n", xi)
	fmt.Printf("alamat x: %p\n", &x)
	updateFighterJet(&x, "F-35C", "hitam", "assault", 342)
	fmt.Println(&x)

	fmt.Println("====================================")
	var t = new(fighterJet)
	fmt.Printf("alamat variable t : %p \n", t)

	fmt.Println("====================================")
	fmt.Println("embeded struct")

	var su25 = fighterJet{
		merek:  "sukhoi",
		warna:  "putih",
		jenis:  "hybrid",
		jumlah: 667,
		buah: buah{
			nama:  "salak",
			hijau: false,
		},
	}
	fmt.Println(su25)
	fmt.Println(su25.buah.hijau)
}

func updateFighterJet(
	newFJ *fighterJet,
	merek, warna, jenis string,
	jumlah int) {
	newFJ.merek = merek
	newFJ.jenis = jenis
	newFJ.warna = warna
	newFJ.jumlah = jumlah

	fmt.Println(newFJ)
}

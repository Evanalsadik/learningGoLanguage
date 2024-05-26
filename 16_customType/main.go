package main

import "fmt"

type Ruangan struct {
	nama  string
	nomor int
	Fahrenheit
}

type Fahrenheit int

func main() {
	var temperature Fahrenheit = 546
	fmt.Println("suhu di tempat ini cukup panas:", temperature)
	temperature.toCelcius()

	temperature.decrease(5)
	fmt.Println("suhu di tempat ini turun sedikit menjadi:", temperature)

	var lobby = Ruangan{"Lobby", 13, 81}
	fmt.Println(lobby)

}

func (f Fahrenheit) toCelcius() {
	hasilKonversi := 0.56 * float64((f - 32))
	fmt.Printf("suhu setelah dikonversi: %.2f\n", hasilKonversi)
}

func (f *Fahrenheit) decrease(value float64) {
	*f -= Fahrenheit(value)
}

package main

import "fmt"

func main() {
	var hargaTotal float64
	var persenDiskon float64

	fmt.Println("=== Ini Kalkulator ===")

	fmt.Println("Masukkan harga barang : ")
	fmt.Scan(&hargaTotal)

	fmt.Println("Masukkan diskon(% : ")
	fmt.Scan(&persenDiskon)

	nominalDiskon := hargaTotal * (persenDiskon / 100)
	hargaAkhir := hargaTotal - nominalDiskon

	fmt.Println("---------------------")
	fmt.Println("Potongan Harga : Rp %.2f\n", nominalDiskon)
	fmt.Println("Total Bayar	: Rp %.2f\n", hargaAkhir)
	fmt.Println("---------------------")
}
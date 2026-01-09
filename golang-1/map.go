package main

import "fmt"

func main() {
	hargaPokok := make(map[string]int)

	hargaPokok["laptop"] = 12000000
	hargaPokok["balok kayu"] = 100000

	fmt.Println("harga laptop", hargaPokok["laptop"])

}
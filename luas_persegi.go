package main

import (
	"fmt"
)

func main() {

	var sisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas := sisi * sisi

	fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, luas)
}

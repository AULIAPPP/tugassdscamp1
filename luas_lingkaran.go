package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	luas := math.Pi * jariJari * jariJari

	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luas)
}

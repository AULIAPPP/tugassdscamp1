package main

import (
	"fmt"
)

func main() {
	var periode float64

	fmt.Print("Masukkan periode (detik): ")
	fmt.Scan(&periode)

	frekuensi := 1.0 / periode

	fmt.Printf("Frekuensi adalah %.2f Hz\n", frekuensi)
}

package main

import (
	"fmt"
)

func main() {
	var massa, tinggi, kecepatan float64

	fmt.Print("Masukkan massa benda (kg): ")
	fmt.Scan(&massa)

	fmt.Print("Masukkan tinggi benda di atas permukaan tanah (m): ")
	fmt.Scan(&tinggi)

	fmt.Print("Masukkan kecepatan benda (m/s): ")
	fmt.Scan(&kecepatan)

	gravitasi := 9.81
	energiPotensial := massa * gravitasi * tinggi

	energiKinetik := 0.5 * massa * kecepatan * kecepatan

	fmt.Printf("Energi Potensial: %.2f joule\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f joule\n", energiKinetik)
}

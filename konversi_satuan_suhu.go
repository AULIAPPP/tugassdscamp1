package main

import (
	"fmt"
)

func main() {
	var suhu float64
	var satuan string

	fmt.Print("Masukkan suhu: ")
	fmt.Scan(&suhu)

	fmt.Print("Satuan suhu (C, F, K): ")
	fmt.Scan(&satuan)

	switch satuan {
	case "C", "c":
		// KCelsius ke Fahrenheit
		fahrenheit := (suhu * 9 / 5) + 32
		// Celsius ke Kelvin
		kelvin := suhu + 273.15

		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", suhu, fahrenheit)
		fmt.Printf("%.2f Celsius = %.2f Kelvin\n", suhu, kelvin)

	case "F", "f":
		// Fahrenheit ke Celsius
		celsius := (suhu - 32) * 5 / 9
		// Fahrenheit ke Kelvin
		kelvin := (suhu + 459.67) * 5 / 9

		fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", suhu, celsius)
		fmt.Printf("%.2f Fahrenheit = %.2f Kelvin\n", suhu, kelvin)

	case "K", "k":
		// Kelvin ke Celsius
		celsius := suhu - 273.15
		// Kelvin ke Fahrenheit
		fahrenheit := (suhu * 9 / 5) - 459.67

		fmt.Printf("%.2f Kelvin = %.2f Celsius\n", suhu, celsius)
		fmt.Printf("%.2f Kelvin = %.2f Fahrenheit\n", suhu, fahrenheit)

	default:
		fmt.Println("Satuan suhu tidak valid. Gunakan C, F, atau K.")
	}
}

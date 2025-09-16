package main

import "fmt"

func main() {
	var angka1, angka2 int
	fmt.Print("Masukkan angka pertama: ")
	// Input angka 1
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan angka kedua: ")
	// Input angka 2
	fmt.Scanln(&angka2)

	// Perbandingan
	fmt.Printf("%d < %d ? %v \n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d > %d ? %v \n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d == %d ? %v \n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v \n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d <= %d ? %v \n", angka1, angka2, angka1 <= angka2)
	fmt.Printf("%d >= %d ? %v \n", angka1, angka2, angka1 >= angka2)
}
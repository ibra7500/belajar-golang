package main

import "fmt"

func main() {
	x := 12
	y := 3

	divided := x / y
	multiply := x * y
	modulus  := x % y

	fmt.Println("Angka yang akan dihitung: ", x, y)
	fmt.Println("Hasil pembagian: ", divided)
	fmt.Println("Hasil perkalian: ", multiply)
	fmt.Println("Hasil modulus: ", modulus)

	x++
	fmt.Println("Hasil increment: ", x)
}
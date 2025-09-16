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
	// fmt.Printf("%d < %d ? %v \n", angka1, angka2, angka1 < angka2)
	// fmt.Printf("%d > %d ? %v \n", angka1, angka2, angka1 > angka2)
	// fmt.Printf("%d == %d ? %v \n", angka1, angka2, angka1 == angka2)
	// fmt.Printf("%d != %d ? %v \n", angka1, angka2, angka1 != angka2)
	// fmt.Printf("%d <= %d ? %v \n", angka1, angka2, angka1 <= angka2)
	// fmt.Printf("%d >= %d ? %v \n", angka1, angka2, angka1 >= angka2)

	// Kalo pake if statement
	if angka1 < angka2 {
		fmt.Println("Angka", angka1 , "lebih kecil daripada angka", angka2)
	} else if angka1 > angka2 {
		fmt.Println("Angka", angka1 , "lebih besar daripada angka", angka2)
	} else if angka1 == angka2 {
		fmt.Println("Angka", angka1 , "sama dengan daripada angka", angka2)
	}

	// shorthand if
	if angka3 := angka1*2; angka3 < 10 {
		fmt.Println("Hasil kurang dari 10: ", angka3)
	}
	
}
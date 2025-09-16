package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Jumlah Elemen di Array Numbers: ", len(numbers)) //hitung jumlah elemen
	fmt.Println("Index ke-2: ", numbers[2])
	
	numbers[1] = 80 //change value di index ke-1
	fmt.Println("Index ke-1: ", numbers[1])  

	// Coba perulangan
	fmt.Println("Ini adalah array: ")
	for index, value := range numbers {
		fmt.Println("Index ke-", index, " = ", value)
	}

	arr1 := [3]int{10, 20, 30}
	arr2 := [3]int{40, 50, 60}

	fmt.Println("arr1 == arr2 ? ", arr1 == arr2) //expected: false
	fmt.Println("arr1 != arr2 ? ", arr1 != arr2) //expected: true
}

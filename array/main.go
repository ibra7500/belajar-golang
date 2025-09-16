package main

import "fmt"

func main() {
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println("Array angka: ", numbers)
	numbers[0] = 50
	fmt.Println("Index pertama: ", numbers[0]) 
	fmt.Println("Index Terakhir: ", numbers[2]) 
	fmt.Println("Panjang Array: ", len(numbers))
}
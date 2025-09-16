package main

import "fmt"

func greeting(nama string, waktu string) {
	fmt.Println("Hello", nama, ", sekarang waktu", waktu)
}

func mathOperation(a, b int) (multiplyResult int) {
	multiplyResult = a * b
	return
}

// coba variadic function
func jumlahAngka(angka ...int) int {
	total := 0
	for _, value := range angka {
		total += value // total = total + value
	}
	return total
}

func main() {
	greeting("Iber", "Pagi")
	fmt.Println(mathOperation(10, 2))
	fmt.Println(jumlahAngka(10, 20, 30))
}

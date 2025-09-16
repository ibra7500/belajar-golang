package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi panic, dan sudah ditangani:", r)
	}
}

func divideNumber(a,b int) {
	defer handlePanic() //handle panic
	fmt.Printf("Divide %d with %d\n", a, b)
	result := a / b //program akan break ketika b = 0 
	fmt.Println("Result:", result)
}

func main() {
	fmt.Println("Mulai Program")
	divideNumber(10, 5)
	divideNumber(5, 0)
	fmt.Println("Program Selesai")
}
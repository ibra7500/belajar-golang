package main

import "fmt"

func main() {
	defer fmt.Println("Defer tetap berjalan sebelum panic")
	fmt.Println("Sebelum panic")
	panic("Ada sesuatu yang fatal")

	// baris setelah panic tidak akan dieksekusi
	fmt.Println("Setelah panic") //ga bakal jalan
}
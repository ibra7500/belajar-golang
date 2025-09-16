package main

import "fmt"

func main() {
	defer fmt.Println("A: Dieksekusi paling akhir")
	fmt.Println("B: Diekseksui paling awal")
}
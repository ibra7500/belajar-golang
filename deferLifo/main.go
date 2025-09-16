package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("Defer ke-1 x=", x)
	x = 2
	defer fmt.Println("Defer ke-2 x=", x)

	fmt.Println("body, x =", x)
}
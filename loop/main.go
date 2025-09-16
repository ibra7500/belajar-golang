package main

import "fmt"

func main() {
	// for standar
	for i := 0; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
	
	// for dengan konsep while
	j := 0
	for j <= 10 {
		fmt.Println("Perulangan konsep while ke-", j)
		j++
	}

	// for infinite
	k := 5
	for {
		fmt.Println("Perulangan konsep infinite ke-", k)
		k++
		if k == 15 {
			break //keluar dari loop
		}
	}

	fruits := [5]string {"Apel", "Jeruk", "Cherry", "Melon", "Semangka"}
	for index, value := range fruits {
		fmt.Println("Item buah ke-", index, " = ", value)
	}
}
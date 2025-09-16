package main

import "fmt"
import "strings"

func main() {

	text := "Hello World"

	fmt.Println("Lowercase: ", strings.ToLower(text))
	fmt.Println("Uppercase: ", strings.ToUpper(text))

	// Cek karakter pertama, case sensitive
	fmt.Println("Starts with Hello?", strings.HasPrefix(text, "Hello"))

	// Cek karakter mengandung kata tertentu atau tidak, case sensitive
	fmt.Println("Contains world?", strings.Contains(text, "world"))

	// Pisahin string
	fruits := "Apple, Banana, Mango, Guava"

	split_fruits := strings.Split(fruits, ",")
	fmt.Println("Split Fruits: ", split_fruits)

	// Ganti bagian string
	newText := strings.ReplaceAll(text, "Hello World", "Test Golang")
	fmt.Println("Replaced Characters: ", newText)
}
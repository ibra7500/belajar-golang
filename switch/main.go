package main

import "fmt"

func main() {
	angka := 1

	switch angka {
	case 1:
		fmt.Println("Angka satu")
	case 2:
		fmt.Println("Angka dua")
	default:
		fmt.Println("Angka tidak dikenal")
	}

	// short switch statement
	switch days := "Senin"; days {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak valid")
	}
}
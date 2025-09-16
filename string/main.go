package main

import "fmt"

func main() {
	nama := "Ibrahim"
	deskripsi := `Test Multi Line Paragraf
		di Golang`

	detail := nama + " " + "Bekasi"

	fmt.Println("Nama: ", nama)
	fmt.Println("Deskripsi: ", deskripsi)

	fmt.Println("Join String:", detail)

	nama_length := len(nama)
	detail_length := len(detail)

	fmt.Println("Panjang karakter nama: ", nama_length)
	fmt.Println("Panjang karakter detail: ", detail_length)
	
	first_nama_byte:= nama[0]
	fmt.Println("Byte pertama nama: ", first_nama_byte)

	first_nama_char:= nama[0]
	fmt.Printf("Char pertama nama: %c \n", first_nama_char)
}
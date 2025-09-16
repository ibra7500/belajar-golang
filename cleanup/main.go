package main

import "fmt"

func cleanup() {
	// anggap cleanup resource
	fmt.Println("Cleanup: membersihkan resources...")
}

func bacaConfig(namaFile string) {
	defer cleanup() //akan terpanggil jika ada panic

	// recover diletakkan dalam defer untuk menangkap panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)	
		}
	}()

	if namaFile == "" {
		panic("Nama file tidak boleh kosong")
	}
}

func main() {
	bacaConfig("")
	fmt.Println("Program tetap berjalan")
}
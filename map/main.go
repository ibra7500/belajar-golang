package main

import "fmt"

func main() {
	siswa := map[string]string{
		"nama": "Test Nama Siswa",
		"kelas": "A",
	}
	fmt.Println("Nama Siswa: ", siswa["nama"])
	fmt.Println("Kelas: ", siswa["kelas"])
	
	siswa["jurusan"] = "IPA"
	fmt.Println("Jurusan: ", siswa["jurusan"])

	mapSiswa := map[string]int{
		"iber": 90,
		"iber lagi": 80,
	}
	fmt.Printf("Panjang Map Siswa: ", len(mapSiswa))
}
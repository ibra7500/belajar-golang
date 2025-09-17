// Belajar struct di go
package main

import "fmt"
type Address struct {
	City string
	Province string
}
type User struct {
	Name  string
	Email string
	Age   int
	Address
}

type PersegiPanjang struct {
	Panjang float64
	Lebar float64
}

func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}


func main() {
	user := User{
		Name:  "Test Nama",
		Email: "testemail@gmail.com",
		Age:   25,
		Address: Address{
			City: "Bekasi",
			Province: "Jawa Barat",
		},
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Age:", user.Age)
	fmt.Println("City:", user.City)
	fmt.Println("Province:", user.Province)

	hitungPersegiPanjang := PersegiPanjang{10.5, 20.12}
	fmt.Println("Hasil Luas Persegi Panjang:", hitungPersegiPanjang.Luas())

}
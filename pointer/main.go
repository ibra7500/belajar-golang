package main

import "fmt"

// func changeName(changedName *string) {
// 	*changedName = "Iber"
// }

func heal(hp *int) {
	*hp += 50
	fmt.Printf("Anda mendapatkan heal sebesar: %d\n", *hp)
}

func damage(hp *int, damage int) {
	*hp -= damage
	fmt.Printf("Anda terkena damage sebesar: %d, sisa HP menjadi: %d\n", damage, *hp)

	if *hp <= 0 {
		fmt.Printf("Game over! HP anda habis. Sisa hp: %d", *hp)
	}
}

func main() {
	// nama := "Test"
	// fmt.Println("Nama awal:", nama)
	// changeName(&nama)
	// fmt.Println("Name: ", nama)

	healthPoint := 100
	damage(&healthPoint, 75)
	heal(&healthPoint)
	damage(&healthPoint, 100)

}
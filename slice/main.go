package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	firstTwoNumbers := numbers[:2] //get array dari index 0(awal) sampai index 1 (2 angka pertama, tidak termasuk index ke-2)
	lastThirdNumbers := numbers[2:] //get array dari index 2 sampai terakhir
	fmt.Println("Array: ", numbers)
	fmt.Println("Dua angka pertama: ", firstTwoNumbers)
	fmt.Println("Tiga angka terakhir: ", lastThirdNumbers)

	// Get array dari index ke-1 sampai ke-3 (4 tidak termasuk)
	sliceNumbers := numbers[1:4]
	fmt.Println("Sliced numbers: ", sliceNumbers)

	// get jumlah data dalam array
	fmt.Println("Jumlah data dalam array setelah dislicing: ", len(sliceNumbers))

	// buat slice  dengan panjang 3, kapasitas 5
	slice1 := make([]int, 3, 5)
	fmt.Printf("Slice: %v, Panjang: %d, Kapasitas: %d \n", slice1, len(slice1), cap(slice1))

	for index, _ := range slice1 {
		slice1[index] = (index + 1) * 2
	}

	slice1 = append(slice1, 10, 20, 30)
	fmt.Printf("Slice Terbaru: %v, Panjang: %d, Kapasitas: %d \n", slice1, len(slice1), cap(slice1))

	// copy slice dari slice1 ke slice3
	slice2 := make([]int, 2)
	slice3 := copy(slice2, slice1)
	fmt.Println("Hasil copy slice2: ", slice2)
	fmt.Println("Jumlah tersalin ke slice3: ", slice3)
}
package main

import "fmt"

func main() {
	var i8 int8 = -64
	var i32 int32 = 100
	var i64 int64 = -1250
	var ui8 uint8 = 64
	var ui32 uint32 = 100
	var ui64 uint64 = 1250

	fmt.Println("Signed Integers")
	fmt.Printf("Nilai %v adalah int8 \n", i8)
	fmt.Printf("Nilai %v adalah int32 \n", i32)
	fmt.Printf("Nilai %v adalah int64 \n", i64)

	fmt.Println("Unsigned Integers")
	fmt.Printf("Nilai %v adalah uint8 \n", ui8)
	fmt.Printf("Nilai %v adalah uint32 \n", ui32)
	fmt.Printf("Nilai %v adalah uint64 \n", ui64)

	// float
	var f32 float32 = 3.141414
	var f64 float64 = 150213.1238990

	fmt.Println("Float Number")
	fmt.Printf("Nilai %v adalah float32 \n", f32)
	fmt.Printf("Nilai %v adalah float64 \n", f64)
}

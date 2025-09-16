package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score int = 90
	var text string = strconv.Itoa(score) //Itoa convert int jadi string, kalo Atoi dari string jadi int

	fmt.Println("Nilai ujian: ", text)
}
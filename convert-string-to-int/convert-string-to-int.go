package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "1000"
	score, err := strconv.Atoi(text) //Itoa convert int jadi string, kalo Atoi dari string jadi int

	if err != nil {
		fmt.Println("Terdapat error: ", err.Error())
	} else {
		fmt.Println("Nilai ujian: ", score)
	}

}
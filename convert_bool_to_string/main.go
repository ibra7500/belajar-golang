package main

import (
	"fmt"
	"strconv"
)

func main() {
	truth := true
	convertTruth := strconv.FormatBool(truth) //boolean -> string

	fmt.Println("Hasil Convert Boolean ke String: ", convertTruth)

	var realita string = "false"
	convertRealita, err := strconv.ParseBool(realita) //string -> boolean

	if err != nil {
		fmt.Println("Error convert boolean ke string: ", err.Error())
	} else {
		fmt.Println("Hasil Convert String ke Boolean: ", convertRealita)
	}

}
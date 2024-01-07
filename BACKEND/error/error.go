package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Tuliskan angka: ")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "ini angka")
	}else {
		fmt.Println(input, "ini bukan angka")
		fmt.Println(err.Error())
	}
}
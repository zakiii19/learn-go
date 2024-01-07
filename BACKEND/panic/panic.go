package main

import (
	"fmt"
	"strings"
	"errors"
)

func validate (input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Tidak bisa kosong")
	}
	return true, nil
}

func main(){
	var name string
	fmt.Print("Tulis namamu: ")
	fmt.Scan(&name)
	if valid, err :=validate(name); valid{
		fmt.Println("halo", name)
	}else{
		panic(err.Error())
		fmt.Println("end")
	}
}
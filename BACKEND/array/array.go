package main

import "fmt"
func main(){
	var fruits = [4]string{"apple", "orange", "banana", "burger"}
	fmt.Println("jumlah elemen \t\t", len(fruits))
	fmt.Println("isi semua elemen \t\t", fruits)
	fmt.Println(fruits[3])
}
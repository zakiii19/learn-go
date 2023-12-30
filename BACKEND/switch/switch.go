package main

import "fmt"
func main(){
	var point = 10
	switch point {
	case 8:
		fmt.Println("Lulus dengan niai sempurna")
	case 5:
		fmt.Println("Lulus")
	case 4:
		fmt.Println("Hampir Lulus")
	default:
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}
}
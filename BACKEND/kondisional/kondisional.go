package main

import "fmt"
func main(){
	var point = 10
	if point == 10{
		fmt.Println("Lulus dengan niai sempurna")
	} else if point > 5{
		fmt.Println("Lulus")
	} else if point == 4{
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}
}
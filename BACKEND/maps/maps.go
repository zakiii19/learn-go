package main

import "fmt"
func main(){
	var chicken map[string]int
	chicken = map [string]int{}
	chicken["senin"] = 56
	chicken["selasa"] = 75
	fmt.Println("januari", chicken["rabu"], chicken["selasa"])
	
}
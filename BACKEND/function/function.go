package main

import (
	"fmt"
	"strings")
func main(){
	var names = []string{"My", "Skill"}
	printMessages("hallo", names)
}

func printMessages(message string, arr []string) {
	var nameString =strings.Join(arr," ")
	fmt.Println(message, nameString)
	
}
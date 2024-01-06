package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("uhuy")
	os.Exit(1)
	fmt.Println("ahahaha")
}
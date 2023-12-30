package main

import "fmt"
type student struct{
	name string
	grade int
}
func main(){
	var s1 student
	s1.name = "zaki"
	s1.grade = 6

	fmt.Println("nama : ", s1.name)
	fmt.Println("kelas : ", s1.grade)

}
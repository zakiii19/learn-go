package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name":"Zaki", "Age":13}`
	var jsonData = []byte(jsonString)
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodeData = data2.(map[string]interface{})
	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age :", decodeData["Age"])
}
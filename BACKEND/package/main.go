package main

import "fmt"
import "backend/package/lib/math"

func main() {
	xs := []float64{1,2,3,4}
	avg := math.average(xs)
	fmt.Println(avg)
}
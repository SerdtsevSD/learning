package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(hypo(a, b))
}
func hypo(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

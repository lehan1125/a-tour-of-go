package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for math.Abs(z*z-x) > 0.00001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

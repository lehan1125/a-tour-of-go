package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't sqrt negative numbers and complex numbers: %f", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for (z*z - x) > 0.0000000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

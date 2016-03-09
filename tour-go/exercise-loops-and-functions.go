package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := float64(1)
	for i := 0; i < 100; i++ {
		z = z - (z*z-x)/2/z
	}
	return z
}
func Sqrt2(x float64) float64 {
	z := float64(1)
	tmp := float64(0)

	for math.Abs(z-tmp) > 0.0000001 {
		tmp = z
		z = z - (z*z-x)/2/z
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt2(2))
	fmt.Println(math.Sqrt(2))
}

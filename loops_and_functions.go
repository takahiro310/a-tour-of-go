package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	zz := z
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		if z == zz {
			break
		}
		zz = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sqrt(3))
	fmt.Println(math.Sqrt(4))
}

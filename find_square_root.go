package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Using this formula to find the square root: z -= (z*z - x) / (2*z)
	z := x/2
	for i := 1; i < 10; i++{
		fmt.Println("Calculating... Current value is: ", z)
		old_value := z
		z = z - ((z * z - x) / (2 * z))
		if old_value == z {
			return z
		}
	}
	return z
}

func main() {
	var x float64 = 25
	fmt.Printf("\nCalculated square root of %v is %v\n", x, Sqrt(x))
	fmt.Printf("(From math.Sqrt lib): Square root of %v is %v", x, math.Sqrt(x))
}

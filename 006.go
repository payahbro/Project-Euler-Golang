package main

import "math"

// NOTE:
// The math formula comes from arithmetic.

func SumSquareDifference(n float64) float64 {
	a := 1.0
	b := 1.0
	res := 0.0

	for i := 1; i <= 10; i++ {
		res = res + math.Pow(float64(i), 2)
	}

	return math.Pow(n*(2*a+(n-1)*b)/2, 2) - res
}

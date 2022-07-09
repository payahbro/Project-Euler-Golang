package main

func LargestPrimeFactor() int {
	var res []int
	target := 600851475143

	for prime := 2; prime <= target; prime++ {
		if target%prime == 0 {
			res = append(res, prime)
			target = target / prime
		}
	}

	_, maxRes := MinMax(res)

	return maxRes
}

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

	_, max := MinMax(res)

	return max
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

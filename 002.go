package main

func EvenFibNumber() int {
	prev := 1
	curr := 2
	res := 0

	for curr < 4000000 {
		if curr%2 == 0 {
			res += curr
		}
		prev, curr = curr, curr+prev
	}

	return res
}

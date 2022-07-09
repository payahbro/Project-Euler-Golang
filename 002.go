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

// Every third fibonacci number is even ex: 1, 1, 2, 3, 5, 8, 13, 21, 34, ......

func OptimEvenFibNumber() int {
	a := 1
	b := 1
	sum := a + b
	res := 0

	for sum < 4000000 {
		res += sum
		a = b + sum
		b = a + sum
		sum = a + b
	}

	return res
}

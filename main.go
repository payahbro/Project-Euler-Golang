package main

import "fmt"

func main() {
	// Day 1 - Multiples of 3 or 5
	fmt.Println("Day 1: ", Multiples())

	// Day 2 - Even Fibonacci numbers
	fmt.Println("Day 2: ", EvenFibNumber())
	fmt.Println("Day 2: ", OptimEvenFibNumber())

	// Day 3 - The Largest Prime Factor
	fmt.Println("Day 3: ", LargestPrimeFactor())

	// Day 4 - The Largest Palindrome Product
	fmt.Println("Day 4: ", LargestPalindromeProduct())

	// Day 5 -  Smallest Multiple
	fmt.Println("Day 5: ", SmallestMultiple())

	// Day 6 - Sum Square Difference
	fmt.Println("Day 6: ", SumSquareDifference(100.0))

	// Day 7 - 10001 Prime
	fmt.Println("Day 7: ", IndexPrime(10001))
}

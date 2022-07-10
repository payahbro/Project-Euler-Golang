package main

func LargestPalindromeProduct() int {

	var palindrome = 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if ReverseInt(i*j) == i*j && i*j > palindrome {
				palindrome = i * j
			}
		}
	}

	return palindrome
}

package main

func LargestPalindromeProduct() int {

	var palindrome []int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if ReverseInt(i*j) == i*j {
				palindrome = append(palindrome, i*j)
			}
		}
	}

	_, maxPalindrome := MinMax(palindrome)

	return maxPalindrome
}

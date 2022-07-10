package main

// NOTE:
// For optimal solutions you can use LCM (Factor Tree) of 1, 2, 3, ....., 20.
// The result of LCM is 16*9*5*7*11*13*17*19 --> 232.792.560 (for prime number you don't need to Prime Factorization)

func SmallestMultiple() int {
	// The result must be more than 2520, so we start from that.
	res := 2520
	validator := 0

	for true {
		for i := 2; i <= 10; i++ {
			if res%i == 0 {
				validator++
				if validator == 9 {
					return res
				}

			} else {
				break
			}
		}

		validator = 0
		res++
	}

	return 0
}

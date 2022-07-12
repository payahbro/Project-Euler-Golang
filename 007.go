package main

func IndexPrime(idx int) int {
	var primeList []int
	n := 0

	for len(primeList) < idx {
		n += 1
		if IsPrime(n) == true {
			primeList = append(primeList, n)
		}
	}

	return primeList[len(primeList)-1]
}

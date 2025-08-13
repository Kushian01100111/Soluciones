package main

func countBits(n int) []int {
	amountOfOnes := make([]int, n+1)
	lastPow := 1
	for i := 1; i <= n; i++ {
		amountOfOnes[i] = 0
		if lastPow*2 == i {
			lastPow = i
		}
		amountOfOnes[i] = amountOfOnes[i-lastPow] + 1
	}

	return amountOfOnes
}

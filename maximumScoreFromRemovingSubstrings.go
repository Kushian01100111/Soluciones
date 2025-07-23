package main

func eliminatePairs(a string, sumPair int, s string, score int) (string, int) {
	stack := []rune{}
	for _, char := range s {
		stack = append(stack, char)
		n := len(stack)
		if n >= 2 && string(stack[n-2:]) == a {
			stack = stack[:n-2]
			score += sumPair
		}
	}
	return string(stack), score
}

func maximumGain(s string, x int, y int) int {
	score := 0
	biggerPair, sumBiggerPair := "", 0
	smallerPair, sumSmallerPair := "", 0

	if x > y {
		biggerPair, sumBiggerPair = "ab", x
		smallerPair, sumSmallerPair = "ba", y
	} else {
		biggerPair, sumBiggerPair = "ba", y
		smallerPair, sumSmallerPair = "ab", x
	}

	stringAfter, gained := eliminatePairs(biggerPair, sumBiggerPair, s, score)
	s, score = eliminatePairs(smallerPair, sumSmallerPair, stringAfter, gained)

	return score
}

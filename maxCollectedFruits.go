package main

func maxCollectedFruits(fruits [][]int) int {
	lenFruits := len(fruits) - 1
	first := 0

	for i := 0; i < len(fruits); i++ {
		curr := fruits[i][i]
		first += curr
	}

	secund := 0
	//movements2th := len(fruits) - 1
	j := len(fruits) - 1
	for i := 0; i < len(fruits); i++ {
		act := fruits[i][j]
		secund += act
		fruits[i][j] = 0
		if j == lenFruits {
			left := fruits[i+1][j-1]
			down := fruits[i+1][j]
			max := max(left, down)
			if max == left {
				j--
			}
		} else {
			left := fruits[i+1][j-1]
			down := fruits[i+1][j]
			right := fruits[i+1][j+1]
			max := max(max(left, down), right)
			if max == left {
				j--
			} else if max == right {
				j++
			}
		}
	}

	return 4
}

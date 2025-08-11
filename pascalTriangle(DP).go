package main

func generate(numRows int) [][]int {
	pascal := make([][]int, numRows)

	n := 0
	for i := 1; i <= numRows; i++ {
		curr := make([]int, i)
		curr[0], curr[i-1] = 1, 1
		if i > 2 {
			for j := 1; j < i-1; j++ {
				curr[j] = pascal[n-1][j-1] + pascal[n-1][j]
			}
		}
		pascal[n] = curr
		n++
	}

	return pascal
}

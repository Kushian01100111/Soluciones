package main

func getRow(rowIndex int) []int {
	pascal := make([][]int, rowIndex+1)

	for i := 0; i < len(pascal); i++ {
		curr := make([]int, i+1)
		curr[0], curr[len(curr)-1] = 1, 1
		if i > 1 {
			for j := 1; j < len(curr)-1; j++ {
				curr[j] = pascal[i-1][j-1] + pascal[i-1][j]
			}
		}
		pascal[i] = curr
	}
	return pascal[rowIndex]
}

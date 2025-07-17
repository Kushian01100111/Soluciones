func max(n int, m int) int {
	if n > m {
		return n
	}
	return m
}

func maximumLength(nums []int, k int) int {
	modValues := make([][]int, k)

	for i := 0; i < k; i++ {
		temp := make([]int, k)
		modValues[i] = temp
	}
	maxLen := 0

	for j := 0; j < len(nums); j++ {
		currMod := nums[j] % k

		for prevMod := 0; prevMod < k; prevMod++ {
			modValues[prevMod][currMod] = modValues[currMod][prevMod] + 1
			maxLen = max(maxLen, modValues[prevMod][currMod])
		}
	}

	return maxLen
}
package main

func totalFruit(fruits []int) int {
	nums := make(map[int]int, 0)
	maxSum := 0
	currSum := 0
	j := 0
	for i := 0; i < len(fruits); i++ {
		curr := fruits[i]
		nums[curr]++
		currSum++

		for len(nums) > 2 {
			temp := fruits[j]
			currSum--
			nums[temp] -= 1
			j++
			if nums[temp] == 0 {
				delete(nums, temp)
			}
		}
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}

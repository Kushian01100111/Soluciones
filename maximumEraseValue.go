package main

func maximumUniqueSubarray(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	uniqueNums := make(map[int]int, 0)

	maxSum := 0
	currSum := 0
	L := 0
	for _, right := range nums {
		for {
			_, ok := uniqueNums[right]
			if !ok {
				break
			}
			curr := nums[L]
			currSum -= curr
			uniqueNums[curr] -= 1
			L++
			if uniqueNums[curr] == 0 {
				delete(uniqueNums, curr)
			}
		}
		currSum += right
		uniqueNums[right] = 1
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

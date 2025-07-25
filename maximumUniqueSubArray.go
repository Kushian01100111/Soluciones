package main

func maxSum(nums []int) int {
	mapa := make(map[int]int, 0)
	uniqueElems := make([]int, 0, len(nums))

	for _, num := range nums {
		_, ok := mapa[num]
		if !ok {
			mapa[num] = 1
			uniqueElems = append(uniqueElems, num)
		}
	}

	onlyPositives := make([]int, 0, len(uniqueElems))
	max := uniqueElems[0]
	currentSum := 0
	for _, num := range uniqueElems {
		if num >= 0 {
			onlyPositives = append(onlyPositives, num)
			currentSum += num
		} else if num > max {
			max = num
		}
	}
	if len(onlyPositives) == 0 {
		return max
	}

	return currentSum
}

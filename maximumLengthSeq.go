func max(n int, m int) int {
	if n > m {
		return n
	}
	return m
}

func maximumLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	odds, evens, alternating := 0, 0, 0
	isOdd := nums[0] % 2

	for _, n := range nums {
		curr := n % 2
		if curr == 0 {
			odds++
			if curr == isOdd {
				alternating++
				isOdd = 1
			}
		} else if curr == 1 {
			evens++
			if curr == isOdd {
				alternating++
				isOdd = 0
			}
		}
	}

	return max(odds, max(evens, alternating))
}
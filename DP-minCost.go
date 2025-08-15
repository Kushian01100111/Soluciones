package main

func minCostClimbingStairs(cost []int) int {
	minCostfrom0 := cost[0]
	for i := 0; i < len(cost); i++ {
		if i+2 < len(cost) {
			jump := cost[i+1]
			jump2 := cost[i+2]
			if jump < jump2 {
				minCostfrom0 += jump
			} else {
				minCostfrom0 += jump2
				i++
			}
		} else {
			break
		}
	}

	minCostfrom1 := cost[1]
	for i := 1; i < len(cost); i++ {
		if i+2 < len(cost) {
			jump := cost[i+1]
			jump2 := cost[i+2]
			if jump < jump2 {
				minCostfrom1 += jump
			} else {
				minCostfrom1 += jump2
				i++
			}
		} else {
			break
		}
	}

	return max(minCostfrom1, minCostfrom0)
}

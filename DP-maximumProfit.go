package main

import "math"

func maxProfit(prices []int) int {
	bestProfit := math.MinInt32

	for i, val := range prices {
		for j := i + 1; j < len(prices); j++ {
			currProfit := prices[j] - val
			if currProfit > bestProfit {
				bestProfit = currProfit
			}
		}
	}
	if bestProfit > 0 {
		return bestProfit
	}
	return 0
}

package main

import "math"

func maxProfit(prices []int) int {
	cheapestDay := math.MaxInt32
	maxProfit := 0
	for _, val := range prices {
		currProfit := val - cheapestDay
		if val < cheapestDay {
			cheapestDay = val
		} else if currProfit > maxProfit {
			maxProfit = currProfit
		}
	}
	return maxProfit
}

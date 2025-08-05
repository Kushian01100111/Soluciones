package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	unplaced := make([]int, 0, len(fruits))

	for i := 0; i < len(fruits); i++ {
		curr := fruits[i]
		lengthBaskets := len(baskets)
		for j := 0; j < len(baskets); j++ {
			currBasket := baskets[j]
			if currBasket >= curr {
				baskets = append(baskets[:j], baskets[(j+1):]...)
				break
			}
		}
		if lengthBaskets == len(baskets) {
			unplaced = append(unplaced, curr)
		}
	}
	return len(unplaced)
}

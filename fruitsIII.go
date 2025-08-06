package main

type Basket struct {
	Val   int
	Index int
}

type Comparator func(a, b Basket) bool

func getComparator(criterion string) Comparator {
	if criterion == "primero" {
		return func(a, b Basket) bool {
			return a.Val < b.Val
		}
	}
	return func(a, b Basket) bool {
		return a.Index < b.Index
	}
}

func lowerBoundVal(buckets []Basket, n int) (int, bool) {
	left, right := 0, len(buckets)-1
	resultIndex := -1

	for left <= right {
		mid := (left + right) / 2
		if buckets[mid].Val >= n {
			resultIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if resultIndex == -1 {
		return -1, true // No hay ningún elemento >= n
	}
	return resultIndex, false
}

func numOfUnplacedFruitsI(fruits []int, baskets []int) int {
	unplaced := make([]int, 0, len(fruits))
	orderedBasket := make([]Basket, len(baskets))

	for i, n := range baskets {
		orderedBasket[i] = Basket{Val: n, Index: i}
	}

	orderedBasket = mergeSortInt(orderedBasket, "segundo")
	orderedBasket = mergeSortInt(orderedBasket, "primero")

	for _, n := range fruits {
		Val, unplace := lowerBoundVal(orderedBasket, n)
		if unplace {
			unplaced = append(unplaced, n)
		} else {
			orderedBasket = append(orderedBasket[:Val], orderedBasket[Val+1:]...)
		}
	}

	return len(unplaced)
}

package main

import "fmt"

func main() {
	baskets := []Basket{
		{Val: 3, Index: 2},
		{Val: 1, Index: 0},
		{Val: 2, Index: 1},
	}

	sortedByVal := mergeSortInt(baskets, "primero")   // ordena por Val
	sortedByIndex := mergeSortInt(baskets, "segundo") // ordena por Index

	fmt.Printf("%v", sortedByVal)
	fmt.Printf("%v", sortedByIndex)
}

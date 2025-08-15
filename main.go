package main

import "fmt"

func main() {
	arr := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	test := minCostClimbingStairs(arr)
	fmt.Printf("%v", test)
}

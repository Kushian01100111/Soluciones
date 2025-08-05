package main

import "fmt"

func main() {
	fruits := []int{4, 2, 5}
	baskets := []int{3, 5, 4}

	res := numOfUnplacedFruits(fruits, baskets)

	fmt.Printf("%v", res)
}

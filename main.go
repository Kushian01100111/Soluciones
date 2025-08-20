package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	test := sortedArrayToBST(arr)
	fmt.Printf("%v", test)
}

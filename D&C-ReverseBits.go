package main

import (
	"math"
	"strconv"
	"strings"
)

func getBits(n int) string {
	lenN := int(math.Log10(float64(n))) + 1 //Math magic: amouth of digits of n
	if lenN > 1 {
		m := lenN / 2
		divisor := int(math.Pow(10, float64(m))) //Math magic
		left := n / divisor
		right := n % divisor
		return (getBits(left) + getBits(right))
	}

	var num strings.Builder

	for n != 0 {
		curr := strconv.Itoa(n % 2)
		n = n / 2
		num.WriteString(curr)
	}

	return num.String()
}

func reverseBits(n int) int {
	return 0
}

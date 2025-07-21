package main

import (
	"strings"
)

func makeFancyString(s string) string {
	cap := len(s)
	var res strings.Builder
	res.Grow(cap)

	var prevStr byte
	currentCount := 0
	for i := 0; i < len(s); i++ {
		curr := s[i]
		if res.Len() < 1 {
			res.WriteByte(curr)
			prevStr = curr
			currentCount = 1
		} else if prevStr != curr {
			res.WriteByte(curr)
			prevStr = curr
			currentCount = 1
		} else if prevStr == curr && currentCount < 2 {
			res.WriteByte(curr)
			currentCount++
		}
	}

	return res.String()
}

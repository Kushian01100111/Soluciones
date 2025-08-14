package main

import "strings"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var resChar strings.Builder
	resChar.Grow(len(s))
	index := 0
	subCurr := s[index]

	for i := 0; i < len(t); i++ {
		val := t[i]
		if val == subCurr {
			resChar.WriteByte(val)
			if index+1 < len(s) {
				index++
				subCurr = s[index]
			} else {
				break
			}
		}
	}

	res := resChar.String()
	return res == s
}

package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	size := len(s)
	forward := true
	var index int
	ls := make([]string, numRows)
	for i := 0; i < size; i++ {
		ls[index] += fmt.Sprintf("%c", s[i])

		if forward {
			index++
		} else {
			index--
		}

		if index == 0 {
			forward = true
		} else if index == numRows-1 {
			forward = false
		}
	}

	return strings.Join(ls, "")
}

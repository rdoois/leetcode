package main

import (
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	var res int
	signal := 1

	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		signal = -1
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			break
		}

		res = (res * 10) + (int(r-'0') * signal)
		if res < -2147483648 {
			return -2147483648
		}

		if res > 2147483647 {
			return 2147483647
		}

	}
	return res
}

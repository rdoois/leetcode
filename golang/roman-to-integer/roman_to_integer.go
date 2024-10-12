package main

import "fmt"

func romanToInt(s string) int {
	nums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	max := len(s) - 1
	var res, act, nxt int
	for i := 0; i < max; i++ {
		act = nums[fmt.Sprintf("%c", s[i])]
		nxt = nums[fmt.Sprintf("%c", s[i+1])]

		if act < nxt {
			res -= act
		} else {
			res += act
		}
	}

	res += nums[fmt.Sprintf("%c", s[max])]
	return res
}

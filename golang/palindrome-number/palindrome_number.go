package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := x
	var res int
	for a > 0 {
		res = (res * 10) + (a % 10)
		a = a / 10
	}

	return x == res
}

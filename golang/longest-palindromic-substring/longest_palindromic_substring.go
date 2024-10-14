package main

func longestPalindrome(s string) string {
	size := len(s)

	if size == 1 {
		return s
	}

	bucket := make(map[string]int)
	for i := 0; i < size-1; i++ {
		for j := i; j < size; j++ {
			sub := s[i : j+1]
			if isPalindrome(sub) {
				bucket[sub] = j + 1 - i
			}
		}
	}

	var res string
	var a int

	for k, v := range bucket {
		if v > a {
			a = v
			res = k
		}
	}

	return res
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true

}

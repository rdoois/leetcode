package main

func lengthOfLongestSubstring(s string) int {
	var res, left, max int
	bucket := make([]int, 256)
	max = len(s)

	for right := 0; right < max; right++ {
		current := s[right]

		if bucket[current] > left {
			left = bucket[current]
		}

		if right-left+1 > res {
			res = right - left + 1
		}

		bucket[current] = right + 1
	}

	return res
}

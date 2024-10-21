package main

func maxArea(height []int) int {
	var res, left, right, min, amount int
	right = len(height) - 1

	for left < right {
		if height[left] < height[right] {
			min = height[left]
		} else {
			min = height[right]
		}

		amount = min * (right - left)
		if amount > res {
			res = amount
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

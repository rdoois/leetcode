package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var res int
	for i, n := range nums {
		res = target - n
		if _, ok := m[res]; ok {
			return []int{m[res], i}
		}
		m[n] = i
	}

	return []int{}
}

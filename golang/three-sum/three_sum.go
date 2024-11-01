package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res = [][]int{}
	size := len(nums)
	var bucket = make(map[int]int)
	var exists = make(map[string]bool)

	for i, n := range nums {
		bucket[n] = i
	}
	if len(bucket) < 2 && nums[0] == 0 {
		return [][]int{{0, 0, 0}}
	}

	var sum, idx int

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			sum = nums[i] + nums[j]
			idx = sum * -1
			if n, ok := bucket[idx]; ok && n != i && n != j {
				order := []int{idx, nums[i], nums[j]}
				sort.Ints(order)
				key := fmt.Sprintf("%d%d%d", order[0], order[1], order[2])
				if !exists[key] {
					exists[key] = true
					res = append(res, order)
				}
			}
		}
	}

	return res
}

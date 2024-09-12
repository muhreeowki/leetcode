package main

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	triplets := [][]int{}

	slices.Sort(nums)

	for a := 0; a < lenNums-2; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b, c := a+1, lenNums-1; b < c; {
			sum := nums[a] + nums[b] + nums[c]
			if sum > 0 {
				c--
			} else if sum < 0 {
				b++
			} else {
				triplets = append(triplets, []int{nums[a], nums[b], nums[c]})
				b++
				for ; b <= lenNums-2 && nums[b] == nums[b-1]; b++ {
				}
			}
		}
	}

	return triplets
}

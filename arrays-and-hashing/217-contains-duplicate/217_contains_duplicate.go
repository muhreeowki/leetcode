package main

func containsDuplicate(nums []int) bool {
	s := map[int]bool{}
	for _, i := range nums {
		s[i] = true
	}

	return len(s) != len(nums)
}

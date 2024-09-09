package arrays

func productExceptSelf(nums []int) []int {
	l := len(nums)
	answer := make([]int, len(nums))

	product := 1
	for i := range nums {
		answer[i] = product
		product *= nums[i]
	}

	product = 1
	for i := l - 1; i >= 0; i-- {
		answer[i] *= product
		product *= nums[i]
	}

	return answer
}

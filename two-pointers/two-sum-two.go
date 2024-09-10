package twopointers

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		switch {
		case numbers[i]+numbers[j] > target:
			j--
		case numbers[i]+numbers[j] < target:
			i++
		default:
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

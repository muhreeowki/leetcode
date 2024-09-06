package arrays

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{}
	frequencyGroups := map[int][]int{}

	for _, num := range nums {
		frequencyMap[num]++
	}

	for key, value := range frequencyMap {
		frequencyGroups[value] = append(frequencyGroups[value], key)
	}

	result := []int{}
	for i := len(nums); i > 0; i-- {
		for _, num := range frequencyGroups[i] {
			if k > 0 {
				result = append(result, num)
				k--
			}
		}
	}

	return result
}

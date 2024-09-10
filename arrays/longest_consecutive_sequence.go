package arrays

func longestConsecutive(nums []int) int {
	numsLen := len(nums)
	numsSet := map[int]bool{}
	longest := 0

	// Populate the set
	for _, num := range nums {
		numsSet[num] = true
	}

	// loop through nums
	for _, num := range nums {
		// If this is the start of a sequence
		if _, ok := numsSet[num-1]; !ok {
			// Count how many numbers follow num
			for i := 1; i < numsLen; i++ {
				// If the sequence ends
				if _, ok := numsSet[num+i]; !ok {
					break
				}
				longest++
			}
		}
	}

	return longest
}

package arrays

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		groups[key] = append(groups[key], str)
	}

	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

package arrays

import "sort"

func isAnagram(s, t string) bool {
	sChars := []byte(s)
	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] > sChars[j]
	})

	tChars := []byte(t)
	sort.Slice(tChars, func(i, j int) bool {
		return tChars[i] > tChars[j]
	})

	return string(tChars) == string(sChars)
}

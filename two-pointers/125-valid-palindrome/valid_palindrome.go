package main

func isPalindrome(s string) bool {
	newStr := []rune{}
	for _, char := range s {
		if char >= 65 && char <= 90 {
			newStr = append(newStr, char+32)
		}
		if char >= 97 && char <= 122 || char >= 48 && char <= 57 {
			newStr = append(newStr, char)
		}
	}

	i, j := 0, len(newStr)-1
	for i <= j {
		if newStr[i] != newStr[j] {
			return false
		}
		j--
		i++
	}

	return true
}

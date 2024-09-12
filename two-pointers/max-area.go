package main

func maxArea(height []int) int {
	arrLen := len(height)
	highest := 0

	l, r := 0, arrLen-1
	for l < r {
		h := height[l]
		if height[l] > height[r] {
			h = height[r]
		}

		area := h * (r - l)
		if area > highest {
			highest = area
		}

		switch h {
		case height[l]:
			l++
		case height[r]:
			r--
		}
	}

	return highest
}

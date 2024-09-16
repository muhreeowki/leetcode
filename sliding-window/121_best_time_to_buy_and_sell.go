package main

func maxProfit(prices []int) int {
	highest, l, r := 0, 0, 1

	for r < len(prices) {
		profit := prices[r] - prices[l]
		if profit > highest {
			highest = profit
		}

		if prices[l] >= prices[r] {
			l = r
		}
		r++
	}

	return highest
}

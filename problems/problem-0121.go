package problems

import (
	"math"
)

func maxProfit(prices []int) int {
	profit, mini := 0, math.MaxInt32

	for _, price := range prices {
		if mini > price {
			mini = price
		}

		if price-mini > profit {
			profit = price - mini
		}
	}

	return profit
}

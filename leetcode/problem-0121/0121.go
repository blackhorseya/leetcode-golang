package problem

func maxProfit(prices []int) int {
	profit, buyFlag := 0, 0

	for i, price := range prices[1:] {
		if price < prices[buyFlag] {
			buyFlag = i + 1
		}

		if price-prices[buyFlag] > profit {
			profit = price - prices[buyFlag]
		}
	}

	return profit
}

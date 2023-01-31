package maxProfile

import "math"

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices)+1)
	for idx := range dp {
		dp[idx] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for idx := 1; idx < len(prices); idx++ {
		dp[idx][0] = max(dp[idx-1][1]+prices[idx], dp[idx-1][0])
		dp[idx][1] = max(dp[idx-1][1], -prices[idx])
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(arrs ...int) int {
	res := math.MinInt64
	for _, num := range arrs {
		if num > res {
			res = num
		}
	}
	return res
}

func maxProfit3(prices []int) int {
	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

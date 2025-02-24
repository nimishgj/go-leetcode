package main

// start of the solution
import "math"

func maxProfit(prices []int) int {
	minimumPrice := math.MaxInt32
	maxPrice := 0
	for _, price := range prices {
		if price < minimumPrice {
			minimumPrice = price
		} else if (price - minimumPrice) > maxPrice {
			maxPrice = price - minimumPrice
		}
	}
	return maxPrice
}

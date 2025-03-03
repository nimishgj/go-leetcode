package main

import "math"

func maxSubArray(nums []int) int {
	sum, maxSoFar := 0, math.MinInt
	for _, n := range nums {
		sum = max(n, sum+n)
		maxSoFar = max(maxSoFar, sum)
	}
	return maxSoFar
}

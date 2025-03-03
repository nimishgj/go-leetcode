package main

func maxProduct(nums []int) int {
	minProductSoFar, maxProductSoFar, result := nums[0], nums[0], nums[0]
	lengthOfNums := len(nums)

	for index := 1; index < lengthOfNums; index++ {
		num := nums[index]
		tempMax := max(num, max(num*maxProductSoFar, num*minProductSoFar))
		minProductSoFar = min(num, min(num*minProductSoFar, num*maxProductSoFar))
		maxProductSoFar = tempMax
		result = max(result, maxProductSoFar)
	}
	return result
}

package main

// start of the solution

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if val, ok := numMap[target-num]; ok {
			return []int{index, val}
		}
		numMap[num] = index
	}
	return []int{}
}

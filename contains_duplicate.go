package main

// start of the solution

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return true
		}
		numsMap[num] = true
	}
	return false
}

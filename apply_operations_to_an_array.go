package main

func applyOperations(nums []int) []int {
	lengthOfNums := len(nums)
	for index := 0; index < lengthOfNums-1; index++ {
		if nums[index] == nums[index+1] {
			nums[index] *= 2
			nums[index+1] = 0
		}
	}
	writeIndex := 0
	for readIndex := 0; readIndex < lengthOfNums; readIndex++ {
		if nums[readIndex] != 0 {
			nums[writeIndex] = nums[readIndex]
			if writeIndex != readIndex {
				nums[readIndex] = 0
			}
			writeIndex++
		}
	}

	return nums
}
